package conversion

import (
	"math"
	"strings"
)

func ConvertToDecimal(input string) (result int) {
	s := strings.TrimSpace(input)

	result = 0

	for i, v := range s {
		if v == '1' {
			result = result + (int)(math.Pow(2, (float64)(len(s)-1-i)))
		}
	}

	return result
}
