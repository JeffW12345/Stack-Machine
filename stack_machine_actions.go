package main

import (
	"errors"
)

func dupActions(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}
	lastNumberPushed := stack[len(stack)-1]
	return append(stack, lastNumberPushed)
}

func popActions(stack []int) ([]int, error) {
	if len(stack) == 0 {
		return stack, errors.New("cannot apply pop operation to an empty stack")
	}
	return stack[:len(stack)-1], nil
}

func sumActions(stack []int) ([]int, error) {
	if len(stack) == 0 {
		return stack, errors.New("SUM operator applied, but no numbers in stack to add up")
	}

	var total int = 0
	for i := 0; i < len(stack); i++ {
		total += stack[i]
	}
	stack = make([]int, 0)
	return append(stack, total), nil
}

func multiplySymbolActions(stack []int) ([]int, error) {
	if len(stack) < 2 {
		return stack, errors.New("multiply operator applied, but fewer than 2 items in stack")
	}

	lastNumberPushed, numberBeforeLastPushed := stack[len(stack)-1], stack[len(stack)-2]
	product := lastNumberPushed * numberBeforeLastPushed
	stack = stack[:len(stack)-2]
	return append(stack, product), nil
}

func minusSymbolActions(stack []int) ([]int, error) {
	if len(stack) < 2 {
		return stack, errors.New("minus operator applied, but fewer than 2 items in stack")
	}

	lastNumberPushed, numberBeforeLastPushed := stack[len(stack)-1], stack[len(stack)-2]
	difference := lastNumberPushed - numberBeforeLastPushed
	if difference < 0 {
		return stack, errors.New("overflow as result of minus operation < 0")
	}
	stack = stack[:len(stack)-2]
	return append(stack, difference), nil
}

func plusSymbolActions(stack []int) ([]int, error) {
	if len(stack) < 2 {
		return stack, errors.New("plus operator applied, but fewer than 2 items in stack")
	}

	lastNumberPushed, numberBeforeLastPushed := stack[len(stack)-1], stack[len(stack)-2]
	sum := lastNumberPushed + numberBeforeLastPushed
	if sum > 50_000 {
		return stack, errors.New("overflow as result of add operation > 50,000")
	}
	stack = stack[:len(stack)-2]
	return append(stack, sum), nil
}