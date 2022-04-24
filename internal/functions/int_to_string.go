package functions

import "strconv"

func IntToString(category int) string {
	strCategory := strconv.Itoa(category)
	if strCategory == "1" {
		strCategory = "purchase"
	} else if strCategory == "0" {
		strCategory = "sale"
	}
	return strCategory
}
