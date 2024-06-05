package main

import (
	"testing"
)

func TestGivenProcessStringGivenValidInput_producesExpectedOutput(t *testing.T) {
	testInputsToExpectedOutputs := map[string]int{
		"99":              99,
		"1 2 +":           3,
		"3 DUP +":         6,
		"DUP 99":          99,
		"3 DUP * 1 +":     10,
		"1 2 3 4 5 SUM":   15,
		"2 5 -":           3,
	}

	for input, want := range testInputsToExpectedOutputs {
		got, err := processInput(input)
		if err != nil {
			t.Fatalf("unexpected error for input %q: %v", input, err)
		}
		if got != want {
			t.Errorf("for input %q got %d, want %d", input, got, want)
		}
	}
}

func TestGivenProcessStringGivenInvalidInput_producesExpectedOutput(t *testing.T) {
	testInputsToExpectedErrorMessage := map[string]string{
		"50000 1 +":           "overflow as result of add operation > 50,000",
		"50001":               "number input 50001 needs to be between 0 and 50,000 inclusive",
		"-1":                  "number input -1 needs to be between 0 and 50,000 inclusive",
		"1.5":                 "invalid input 1.5: strconv.Atoi: parsing \"1.5\": invalid syntax",
		"XXX-INVALID 1 2 +":   "invalid input XXX-INVALID: strconv.Atoi: parsing \"XXX-INVALID\": invalid syntax",
		"1 2":				   "there should be just one element in the stack at the end of the operation but there are 2 elements",
		"POP":				   "cannot apply pop operation to an empty stack",
		"SUM":				   "SUM operator applied, but no numbers in stack to add up",
		"1 *":				   "multiply operator applied, but fewer than 2 items in stack",
		"*":				   "multiply operator applied, but fewer than 2 items in stack",
		"1 -":				   "minus operator applied, but fewer than 2 items in stack",
		"-":				   "minus operator applied, but fewer than 2 items in stack",
		"2 1 -":			 	"overflow as result of minus operation < 0",
		"1 +":				   "plus operator applied, but fewer than 2 items in stack",
		"+":				   "plus operator applied, but fewer than 2 items in stack",
	}

	for input, want := range testInputsToExpectedErrorMessage {
		_, err := processInput(input)
		if err == nil {
			t.Fatalf("expected error for input %q but got none", input)
		}
		if err.Error() != want {
			t.Errorf("for input %q got error message %q, want error message %q", input, err.Error(), want)
		}
	}
}
