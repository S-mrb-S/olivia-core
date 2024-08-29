package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/MehraB832/olivia_core/locales"

	"github.com/MehraB832/olivia_core/modules/start"

	"github.com/MehraB832/olivia_core/analysis"
	"github.com/MehraB832/olivia_core/user"
	"github.com/MehraB832/olivia_core/util"
	"github.com/gookit/color"
	"github.com/gorilla/websocket"
)

// upgrader -> websocketUpgrader
// websocketUpgrader configures the websocket upgrader for handling connections
var websocketUpgrader = websocket.Upgrader{ // upgrader -> websocketUpgrader
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// RequestMessage -> clientRequestMessage
// RequestMessage is the structure that uses entry connections to chat with the websocket
type clientRequestMessage struct { // RequestMessage -> clientRequestMessage
	Type        int              `json:"type"` // 0 for handshakes and 1 for messages
	Content     string           `json:"content"`
	Token       string           `json:"user_token"`
	Locale      string           `json:"locale"`
	Information user.UserProfile `json:"information"`
}

// ResponseMessage -> serverResponseMessage
// ResponseMessage is the structure used to reply to the user through the websocket
type serverResponseMessage struct { // ResponseMessage -> serverResponseMessage
	Content     string           `json:"content"`
	Tag         string           `json:"tag"`
	Information user.UserProfile `json:"information"`
}

// SocketHandle -> HandleWebSocketConnection
// HandleWebSocketConnection manages the entry connections and replies with the neural network
func HandleWebSocketConnection(w http.ResponseWriter, r *http.Request) { // SocketHandle -> HandleWebSocketConnection
	conn, _ := websocketUpgrader.Upgrade(w, r, nil) // upgrader -> websocketUpgrader
	fmt.Println(color.FgGreen.Render("A new connection has been opened"))

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Unmarshal the json content of the message
		var request clientRequestMessage // RequestMessage -> clientRequestMessage
		if err = json.Unmarshal(msg, &request); err != nil {
			continue
		}

		// Set the information from the client into the cache
		if reflect.DeepEqual(user.RetrieveUserProfile(request.Token), user.UserProfile{}) {
			user.StoreUserProfile(request.Token, request.Information)
		}

		// If the type of requests is a handshake then execute the start modules
		if request.Type == 0 {
			start.ExecuteModules(request.Token, request.Locale)

			message := start.GetMessage()
			if message != "" {
				// Generate the response to send to the user
				response := serverResponseMessage{ // ResponseMessage -> serverResponseMessage
					Content:     message,
					Tag:         "start module",
					Information: user.RetrieveUserProfile(request.Token),
				}

				bytes, err := json.Marshal(response)
				if err != nil {
					panic(err)
				}

				if err = conn.WriteMessage(msgType, bytes); err != nil {
					continue
				}
			}

			continue
		}

		// Write message back to browser
		response := generateReply(request) // Reply -> generateReply
		if err = conn.WriteMessage(msgType, response); err != nil {
			continue
		}
	}
}

// Reply -> generateReply
// generateReply takes the entry message and returns an array of bytes for the answer
func generateReply(request clientRequestMessage) []byte { // Reply -> generateReply, RequestMessage -> clientRequestMessage
	var responseSentence, responseTag string

	// Send a message from res/datasets/messages.json if it is too long
	if len(request.Content) > 500 {
		responseTag = "too long"
		responseSentence = util.SelectRandomMessage(request.Locale, responseTag) // Keeping SelectRandomMessage as is
	} else {
		// If the given locale is not supported yet, set english
		locale := request.Locale
		if !locales.Exists(locale) { // Keeping Exists as is
			locale = "en"
		}

		responseTag, responseSentence = analysis.NewSentence(
			locale, request.Content,
		).Calculate(*cacheInstance, globalNeuralNetworks[locale], request.Token) // Keeping NewSentence and Calculate as is
	}

	// Marshall the response in json
	response := serverResponseMessage{ // ResponseMessage -> serverResponseMessage
		Content:     responseSentence,
		Tag:         responseTag,
		Information: user.RetrieveUserProfile(request.Token),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return bytes
}
