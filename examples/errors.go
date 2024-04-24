package examples

import (
	"errors"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func DemoErrors() {
	demoStringValidation()
	demoIntValidation()
	demoAsfunction()
	demoPanic()
}

func demoStringValidation() {
	fmt.Println(newline + fname(demoStringValidation))
	values := []string{
		"JOHN DOW123",
		"elon mask",
	}
	fmt.Println("values:", values)
	fmt.Println("\nIS_UPPER VALIDATION")
	runValidator(values, validateIsUpper)
	fmt.Println("\nIS_LOWER VALIDATION")
	runValidator(values, validateIsLower)
	const border = 9
	fmt.Printf("\nLENGTH VALIDATION border = %v\n", border)
	runValidator(values, getLengthValidator(9))
}

func demoIntValidation() {
	fmt.Println(newline + fname(demoIntValidation))
	values := make([]int, 10)
	for i := range values {
		values[i] = rint(100)
	}
	fmt.Println("values:", values)

	const border = 50
	fmt.Printf("\nLT_VALIDATION border = %v\n", border)
	runValidator(values, ltValidator(border))
	fmt.Printf("\nGTE_VALIDATION border = %v\n", border)
	runValidator(values, gteValidator(border))
}

func runValidator[T any](values []T, validator func(v T) (T, error)) {
	for _, v := range values {
		v, err := validator(v)

		if err != nil {
			fmt.Printf("Incorrect -- (%s)\n", err)
		} else {
			fmt.Println("Correct --", v)
		}
	}
}

type ValidationError struct {
	code    string
	message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Code: %v, Message: %v", e.code, e.message)
}

func isUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func validateIsUpper(s string) (string, error) {
	if !isUpper(s) {
		return s, ValidationError{code: "UPPER_ERROR", message: fmt.Sprintf("Value is not upper (%v)", s)}
	}
	return s, nil
}

func validateIsLower(s string) (string, error) {
	if !isLower(s) {
		return s, ValidationError{code: "LOWER_ERROR", message: fmt.Sprintf("Value is not lower (%v)", s)}
	}
	return s, nil
}

func getLengthValidator(value int) func(string) (string, error) {
	return func(s string) (string, error) {
		if utf8.RuneCountInString(s) > value {
			return s, ValidationError{code: "LENGTH_ERROR", message: fmt.Sprintf("Value(%v) is longer than %v", s, value)}
		}
		return s, nil
	}
}

func ltValidator(value int) func(int) (int, error) {
	return func(n int) (int, error) {
		if !(n < value) {
			return n, ValidationError{code: "LT_ERROR", message: fmt.Sprintf("Value(%v) is not less than %v", n, value)}
		}
		return n, nil
	}
}

func gteValidator(value int) func(int) (int, error) {
	return func(n int) (int, error) {
		if !(n >= value) {
			return n, ValidationError{code: "GTE_ERROR", message: fmt.Sprintf("Value(%v) is not greater or equal %v", n, value)}
		}
		return n, nil
	}
}

func demoAsfunction() {
	_, err := os.Open("non-existing.txt")
	if err != nil {
		err = fmt.Errorf("Error opening: %w", err)
	}

	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("Using Assert: Error e is of type path error. Error: %v\n", e)
	} else {
		fmt.Println("Using Assert: Error not of type path error")
	}

	var pathError *os.PathError
	fmt.Println(pathError)
	if errors.As(err, &pathError) {
		fmt.Printf("Using As function: Error e is of type path error. Error: %v\n", pathError)
	}
}

type IndexOutOfRange struct {
	index int
	len   int
	Err   error
}

func (e IndexOutOfRange) Error() string {
	return fmt.Sprintf("index: %v, len: %v", e.index, e.len)
}

func getByIndex(slice []string, index int) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	return slice[index], err
}

func demoPanic() {
	slice := []string{"One", "Two"}
	value, err := getByIndex(slice, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
	fmt.Println("End of func")
}
