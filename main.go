package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MehraB832/olivia_core/locales"
	"github.com/MehraB832/olivia_core/training"

	"github.com/MehraB832/olivia_core/dashboard"

	"github.com/MehraB832/olivia_core/util"

	"github.com/gookit/color"

	"github.com/MehraB832/olivia_core/network"

	"github.com/MehraB832/olivia_core/server"
)

var neuralNetworksMapContainer = map[string]network.Network{}

func main() {
	serverPortArg := flag.String("port", "8080", "The port for the API and WebSocket.")
	localeRetrainArg := flag.String("re-train", "", "The locale(s) to re-train.")
	flag.Parse()

	// If the localeRetrainArg isn't empty then retrain the given models
	if *localeRetrainArg != "" {
		executeModelRetraining(*localeRetrainArg)
	}

	// Print the Olivia ASCII text
	oliviaASCIIBanner := string(util.FetchFileContent("res/olivia-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(oliviaASCIIBanner))

	// Create the authentication token
	dashboard.Authenticate()

	for _, individualLocale := range locales.Locales {
		util.GenerateSerializedMessages(individualLocale.Tag)

		neuralNetworksMapContainer[individualLocale.Tag] = training.CreateNeuralNetwork(
			individualLocale.Tag,
			false,
		)
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*serverPortArg = os.Getenv("PORT")
	}

	// Serves the server
	server.StartServer(neuralNetworksMapContainer, *serverPortArg)
}

// executeModelRetraining retrains the given locales
func executeModelRetraining(localeRetrainList string) {
	// Iterate locales by separating them by comma
	for _, individualLocale := range strings.Split(localeRetrainList, ",") {
		trainingFilePath := fmt.Sprintf("res/locales/%s/training.json", individualLocale)
		deleteError := os.Remove(trainingFilePath)

		if deleteError != nil {
			fmt.Printf("Cannot re-train %s model.", individualLocale)
			return
		}
	}
}
