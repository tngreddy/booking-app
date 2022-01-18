package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTkts uint, remainingTkts uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTktNo := userTkts > 0 && userTkts <= remainingTkts
	return isValidName, isValidEmail, isValidTktNo
}
