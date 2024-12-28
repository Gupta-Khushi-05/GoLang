package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var valFloat []float64
	for _, valString := range strings {
		float, err := strconv.ParseFloat(valString, 64)
		if err != nil {
			return nil, errors.New("Failed to convert strings to floats..!")
		}
    valFloat = append(valFloat, float)
	}
	return valFloat, nil
}