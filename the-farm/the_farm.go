package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, count int) (float64, error) {
	totalFodder, err1 := calc.FodderAmount(count)
	fatteningFactor, err2 := calc.FatteningFactor()

	if err1 != nil {
		return 0.0, err1
	}
	if err2 != nil {
		return 0.0, err2
	}

	return totalFodder * fatteningFactor / float64(count), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, count int) (float64, error) {
	if count > 0 {
		return DivideFood(calc, count)
	} else {
		return 0.0, errors.New("invalid number of cows")
	}
}

func ValidateNumberOfCows(count int) error {
	if count < 0 {
		return &InvalidCowsError{
			count:   count,
			message: "there are no negative cows",
		}
	} else if count == 0 {
		return &InvalidCowsError{
			count:   count,
			message: "no cows don't need food",
		}
	} else {
		return nil
	}
}

type InvalidCowsError struct {
	count   int
	message string
}

func (error *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", error.count, error.message)
}
