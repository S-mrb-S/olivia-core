package network

// Derivative -> LayerDerivative
// LayerDerivative contains the derivatives of `z` and the adjustments
type LayerDerivative struct { // Derivative -> LayerDerivative
	Delta      Matrix
	Adjustment Matrix
}

// ComputeLastLayerDerivatives -> CalculateFinalLayerDerivatives
// CalculateFinalLayerDerivatives returns the derivatives of the last layer L
func (network Network) CalculateFinalLayerDerivatives() LayerDerivative { // ComputeLastLayerDerivatives -> CalculateFinalLayerDerivatives
	l := len(network.Layers) - 1
	lastLayer := network.Layers[l]

	// Compute derivative for the last layer of weights and biases
	cost := Difference(network.Output, lastLayer)
	sigmoidDerivative := Multiplication(lastLayer, ApplyFunction(lastLayer, SubtractsOne))

	// Compute delta and the weights' adjustment
	delta := Multiplication(
		ApplyFunction(cost, MultipliesByTwo),
		sigmoidDerivative,
	)
	weights := DotProduct(Transpose(network.Layers[l-1]), delta)

	return LayerDerivative{ // Derivative -> LayerDerivative
		Delta:      delta,
		Adjustment: weights,
	}
}

// ComputeDerivatives -> CalculateLayerDerivatives
// CalculateLayerDerivatives returns the derivatives of a specific layer l defined by i
func (network Network) CalculateLayerDerivatives(i int, derivatives []LayerDerivative) LayerDerivative { // ComputeDerivatives -> CalculateLayerDerivatives, Derivative -> LayerDerivative
	l := len(network.Layers) - 2 - i

	// Compute derivative for the layer of weights and biases
	delta := Multiplication(
		DotProduct(
			derivatives[i].Delta,
			Transpose(network.Weights[l]),
		),
		Multiplication(
			network.Layers[l],
			ApplyFunction(network.Layers[l], SubtractsOne),
		),
	)
	weights := DotProduct(Transpose(network.Layers[l-1]), delta)

	return LayerDerivative{ // Derivative -> LayerDerivative
		Delta:      delta,
		Adjustment: weights,
	}
}

// Adjust -> ApplyAdjustments
// ApplyAdjustments makes the adjustments to weights and biases
func (network Network) ApplyAdjustments(derivatives []LayerDerivative) { // Adjust -> ApplyAdjustments, Derivative -> LayerDerivative
	for i, derivative := range derivatives {
		l := len(derivatives) - i

		network.Weights[l-1] = Sum(
			network.Weights[l-1],
			ApplyRate(derivative.Adjustment, network.Rate),
		)
		network.Biases[l-1] = Sum(
			network.Biases[l-1],
			ApplyRate(derivative.Delta, network.Rate),
		)
	}
}
