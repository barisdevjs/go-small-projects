package helper

import s "strings"

func ValidateUserInput(name string, last string, email string, uTickets uint, rTickets uint) (bool, bool, bool) {
	isValidName := len(name) >= 2 && len(last) >= 2
	isValidEmail := s.Contains(email, "@")
	isValidTicketNumber := uTickets > 0 && uTickets <= rTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
