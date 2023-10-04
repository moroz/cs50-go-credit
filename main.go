package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const CREDIT_CARD_NUMBER = "4003600000000014"

type CreditCardType string

const (
	Visa       CreditCardType = "VISA"
	Mastercard CreditCardType = "MASTERCARD"
	Amex       CreditCardType = "AMEX"
	Invalid    CreditCardType = "INVALID"
)

func validateChecksum(digits []int) bool {
	sum := sumDigitsFromSecondToLast(digits) + sumRemainingDigits(digits)
	return sum%10 == 0
}

func sumDigitsFromSecondToLast(digits []int) int {
	sum := 0
	// count every other digit from second to last
	for i := len(digits) - 2; i >= 0; i -= 2 {
		// double this digit and sum the digits if necessary
		doubledDigit := digits[i] * 2
		if doubledDigit >= 10 {
			sum += (doubledDigit % 10) + 1
		} else {
			sum += doubledDigit
		}
	}
	return sum
}

func sumRemainingDigits(digits []int) int {
	sum := 0
	for i := len(digits) - 1; i >= 0; i -= 2 {
		sum += digits[i]
	}
	return sum
}

func cardNumberToDigits(cardNumber string) ([]int, error) {
	digits := make([]int, len(cardNumber))
	for i, chr := range cardNumber {
		if chr < 48 || chr > 57 {
			msg := fmt.Sprintf("Invalid digit at index %d: %c", i, chr)
			return []int{}, errors.New(msg)
		}
		digits[i] = int(chr - 48)
	}
	return digits, nil
}

func validateChecksumFromString(cardNumber string) bool {
	digits, err := cardNumberToDigits(cardNumber)
	if err != nil {
		return false
	}
	return validateChecksum(digits)
}

func guessCardType(cardNumber string) CreditCardType {
	digits, err := cardNumberToDigits(cardNumber)
	if err != nil || !validateChecksum(digits) {
		return Invalid
	}
	if strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37") {
		return Amex
	}
	if strings.HasPrefix(cardNumber, "4") {
		return Visa
	}
	return Mastercard
}

func main() {
	for {
		var cardNumber string
		fmt.Print("Number: ")
		fmt.Scanln(&cardNumber)

		digits, err := cardNumberToDigits(cardNumber)
		if err != nil {
			continue
		}

		result := validateChecksum(digits)
		fmt.Println(result)
		os.Exit(0)
	}
}
