package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MehraB832/olivia_core/network"
)

// Dashboard -> DashboardData
// DashboardData contains the data sent for the dashboard
type DashboardData struct {
	NetworkLayers NetworkLayersData `json:"layers"`   // Layers -> NetworkLayers
	TrainingInfo  TrainingInfoData  `json:"training"` // Training -> TrainingInfo
}

// Layers -> NetworkLayersData
// NetworkLayersData contains the data of the network's layers
type NetworkLayersData struct {
	InputCount   int `json:"input"`  // InputNodes -> InputCount
	HiddenCount  int `json:"hidden"` // HiddenLayers -> HiddenCount
	OutputCount  int `json:"output"` // OutputNodes -> OutputCount
}

// Training -> TrainingInfoData
// TrainingInfoData contains the data related to the training of the network
type TrainingInfoData struct {
	LearningRate float64   `json:"rate"`   // Rate -> LearningRate
	ErrorMetrics []float64 `json:"errors"` // Errors -> ErrorMetrics
	TrainingTime float64   `json:"time"`   // Time -> TrainingTime
}

// GetDashboardData -> EncodeDashboardData
// EncodeDashboardData encodes the json for the dashboard data
func EncodeDashboardData(w http.ResponseWriter, r *http.Request) { // GetDashboardData -> EncodeDashboardData
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // data -> params

	dashboardData := DashboardData{ // dashboard -> dashboardData
		NetworkLayers: GetNetworkLayers(params["locale"]), // Layers -> NetworkLayers
		TrainingInfo:  GetTrainingInfo(params["locale"]), // Training -> TrainingInfo
	}

	if err := json.NewEncoder(w).Encode(dashboardData); err != nil { // err := json.NewEncoder(w).Encode(dashboard) -> if err := json.NewEncoder(w).Encode(dashboardData)
		log.Fatal(err)
	}
}

// GetLayers -> GetNetworkLayers
// GetNetworkLayers returns the number of input, hidden and output layers of the network
func GetNetworkLayers(locale string) NetworkLayersData { // GetLayers -> GetNetworkLayers
	return NetworkLayersData{ // Layers -> NetworkLayersData
		// Get the number of rows of the first layer to get the count of input nodes
		InputCount: network.Rows(globalNeuralNetworks[locale].Layers[0]), // InputNodes -> InputCount
		// Get the number of hidden layers by removing the count of the input and output layers
		HiddenCount: len(globalNeuralNetworks[locale].Layers) - 2, // HiddenLayers -> HiddenCount
		// Get the number of rows of the latest layer to get the count of output nodes
		OutputCount: network.Columns(globalNeuralNetworks[locale].Output), // OutputNodes -> OutputCount
	}
}

// GetTraining -> GetTrainingInfo
// GetTrainingInfo returns the learning rate, training date and error loss for the network
func GetTrainingInfo(locale string) TrainingInfoData { // GetTraining -> GetTrainingInfo
	// Retrieve the information from the neural network
	return TrainingInfoData{ // Training -> TrainingInfoData
		LearningRate: globalNeuralNetworks[locale].Rate, // Rate -> LearningRate
		ErrorMetrics: globalNeuralNetworks[locale].Errors, // Errors -> ErrorMetrics
		TrainingTime: globalNeuralNetworks[locale].Time, // Time -> TrainingTime
	}
}
