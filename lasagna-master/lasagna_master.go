package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodleLayerCount := 0
	sauceLayerCount := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayerCount++
		}
		if layer == "sauce" {
			sauceLayerCount++
		}
	}

	noodles = noodleLayerCount * 50
	sauce = sauceLayerCount * 0.2
	return
}

func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

func ScaleRecipe(inputQuantities []float64, portions int) []float64 {
	var quantities = make([]float64, len(inputQuantities))
	copy(quantities, inputQuantities)
	for index, quantity := range quantities {
		quantities[index] = quantity / 2 * float64(portions)
	}
	return quantities
}
