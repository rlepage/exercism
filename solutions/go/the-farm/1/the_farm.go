package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message    string
	details    string
	nb_of_cows int
}

func (err SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.nb_of_cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && amount >= 0 {
		amount *= 2
	} else if amount < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	} else if err != nil {
		return 0, err
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, &SillyNephewError{"", "", cows}
	}

	return amount / float64(cows), nil
}
