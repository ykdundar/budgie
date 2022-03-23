package functions

func ConvertBoolToInt(b bool) int {
	value := 0

	if b == true {
		value = 1
	}

	return value
}
