package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTimePerLayer int) int {
	if avgPreparationTimePerLayer == 0 {
		avgPreparationTimePerLayer = 2
	}
	return len(layers) * avgPreparationTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQty []float64 = make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		newQty[i] = quantities[i] * float64(portions) / 2
	}
	return newQty
}