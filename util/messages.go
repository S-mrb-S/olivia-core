package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Message -> DataPacket
// DataPacket contains the message's tag and its contained matched sentences
type DataPacket struct {
	// Tag -> Label
	Label string `json:"tag"`
	// Messages -> Content
	Content []string `json:"messages"`
}

var cachedDataStore = map[string][]DataPacket{}

// SerializeMessages -> GenerateSerializedMessages
// GenerateSerializedMessages serializes the content of `res/datasets/messages.json` in JSON
func GenerateSerializedMessages(region string) []DataPacket {
	var parsedData []DataPacket
	deserializationError := json.Unmarshal(FetchFileContent("res/locales/"+region+"/messages.json"), &parsedData)
	if deserializationError != nil {
		fmt.Println(deserializationError)
	}

	cachedDataStore[region] = parsedData

	return parsedData
}

// GetMessages -> RetrieveCachedMessages
// RetrieveCachedMessages returns the cached messages for the given locale
func RetrieveCachedMessages(region string) []DataPacket {
	return cachedDataStore[region]
}

// GetMessageByTag -> FindMessageByLabel
// FindMessageByLabel returns a message found by the given tag and locale
func FindMessageByLabel(identifier, region string) DataPacket {
	for _, item := range cachedDataStore[region] {
		if identifier != item.Label {
			continue
		}

		return item
	}

	return DataPacket{}
}

// GetMessage -> SelectRandomMessage
// SelectRandomMessage retrieves a message tag and returns a random message chosen from res/datasets/messages.json
func SelectRandomMessage(region, identifier string) string {
	for _, item := range cachedDataStore[region] {
		// Find the message with the right tag
		if item.Label != identifier {
			continue
		}

		// Returns the only element if there aren't more
		if len(item.Content) == 1 {
			return item.Content[0]
		}

		// Returns a random sentence
		rand.Seed(time.Now().UnixNano())
		return item.Content[rand.Intn(len(item.Content))]
	}

	return ""
}
