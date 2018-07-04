package main

import (
	"fmt"
	"strings"
)

const dummyData = "student-email.liv.aC.uk"

func main() {
	fmt.Println(verifyEmail(dummyData))
}

func verifyEmail(email string) bool {
	studentEmailType := []string{"ac.uk", ".edu"}
	email = strings.ToLower(email)

	for _, emailType := range studentEmailType {
		verifyEmail := strings.Contains(email, emailType)
		if verifyEmail {
			return true
		}
	}
	return false
}
