package registration

import (
	"BANK/storage/postgres"
	"errors"
	"fmt"
	"strings"
)

func Register() int {
	var name, email, password string

	fmt.Print("\nEnter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	errs := ValidateUser(name, email)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
		fmt.Println("Validation error. Please try again.")
		return Register()
	}
	postgres.InsertTable(name, email, password)

	fmt.Println("Registration successful.")
	return postgres.Select(name, email, password)
}

func ValidateUser(name, email string) []error {
	var errs []error

	if len(name) < 6 {
		errs = append(errs, errors.New("Name: length cannot be less than 6 characters"))
	}

	if email == "" {
		errs = append(errs, errors.New("Email: cannot be empty"))
	} else {
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			errs = append(errs, errors.New("Email: invalid format"))
		}
	}

	return errs
}
