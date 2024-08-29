package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MehraB832/olivia_core/analysis"

	"github.com/MehraB832/olivia_core/training"

	"github.com/MehraB832/olivia_core/dashboard"

	"github.com/MehraB832/olivia_core/modules/spotify"

	"github.com/MehraB832/olivia_core/network"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	gocache "github.com/patrickmn/go-cache"
)

var (
	// neuralNetworks -> globalNeuralNetworks
	// globalNeuralNetworks is a map to hold the neural network instances
	globalNeuralNetworks map[string]network.Network

	// cache -> cacheInstance
	// cacheInstance initializes the cache with a 5-minute lifetime
	cacheInstance = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve -> StartServer
// StartServer initializes the server with the given neural networks and port
func StartServer(neuralNetworkInstances map[string]network.Network, serverPort string) { // Serve -> StartServer, _neuralNetworks -> neuralNetworkInstances, port -> serverPort
	// Set the current global network as a global variable
	globalNeuralNetworks = neuralNetworkInstances

	// Initializes the router
	router := mux.NewRouter()
	router.HandleFunc("/callback", spotify.CompleteAuth)
	// Serve the websocket
	router.HandleFunc("/websocket", HandleWebSocketConnection)
	// Serve the API
	router.HandleFunc("/api/{locale}/dashboard", EncodeDashboardData).Methods("GET")
	router.HandleFunc("/api/{locale}/intent", dashboard.CreateIntent).Methods("POST")
	router.HandleFunc("/api/{locale}/intent", dashboard.DeleteIntent).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/{locale}/train", TrainNeuralNetwork).Methods("POST") // Train -> TrainNeuralNetwork
	router.HandleFunc("/api/{locale}/intents", dashboard.GetIntents).Methods("GET")
	router.HandleFunc("/api/coverage", analysis.GetCoverage).Methods("GET")

	magentaColor := color.FgMagenta.Render
	fmt.Printf("\nServer listening on the port %s...\n", magentaColor(serverPort)) // magenta -> magentaColor

	// Serves the chat
	if err := http.ListenAndServe(":"+serverPort, router); err != nil { // err := http.ListenAndServe(":"+port, router) -> if err := http.ListenAndServe(":"+serverPort, router)
		panic(err)
	}
}

// Train -> TrainNeuralNetwork
// TrainNeuralNetwork is the route to re-train the neural network
func TrainNeuralNetwork(w http.ResponseWriter, r *http.Request) { // Train -> TrainNeuralNetwork
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Olivia-Token")
	if !dashboard.ChecksToken(token) {
		json.NewEncoder(w).Encode(dashboard.Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	magentaColor := color.FgMagenta.Render
	fmt.Printf("\nRe-training the %s..\n", magentaColor("neural network")) // magenta -> magentaColor

	for locale := range globalNeuralNetworks { // neuralNetworks -> globalNeuralNetworks
		globalNeuralNetworks[locale] = training.CreateNeuralNetwork(locale, true)
	}
}
