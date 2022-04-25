package functions

import "fmt"

func CropNumbers(number float64) string {
	croppedNumber := fmt.Sprintf("%.2f \n", number)

	return croppedNumber
}
