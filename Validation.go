package main

import (
	"errors"
	"regexp"
)

func AgeValidation(age int) error {
	if age <= 0 || age >= 110 {
		return errors.New("invalid Age")
	}
	return nil
}

func EmailValidation(email string) error {
	re, err := regexp.Compile(`^[a-zA-Z0-9._+%-]*@[a-zA-Z0-9._]*\.[a-zA-Z]{2,}$`)
	if err != nil {
		return err
	}
	isEmail := re.MatchString(email)
	if isEmail {
		return nil
	} else {
		return errors.New("invalid Email")
	}
}
