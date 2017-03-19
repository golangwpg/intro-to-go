package circle

import "errors" // HL

const Pi = 3.14

func Area(radius float64) (float64, error) { // HL
	if radius < 0 {
		return 0, errors.New("Invalid radius") // HL
	}

	return Pi * radius * radius, nil
}
