package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	result float64
}

func (c *Calculator) add(num float64) {
	c.result += num
}

func (c *Calculator) subtract(num float64) {
	c.result -= num
}

func (c *Calculator) multiply(num float64) {
	c.result *= num
}

func (c *Calculator) divide(num float64) error {
	if num == 0 {
		return errors.New("can't divide by zero")
	}

	c.result /= num

	return nil
}

func (c *Calculator) reset() {
	c.result = 0
}

func (c *Calculator) getResult() float64 {
	return c.result
}

func (c *Calculator) processInput(input string) error {
	if input == "" {
		return errors.New("input is required")
	}

	parts := strings.Fields(input)

	if len(parts) > 2 {
		return errors.New("input should be [action] [number]")
	}

	action := parts[0]
	var num float64
	var err error

	if (action == "reset" || action == "quit") && len(parts) != 1 {
		return errors.New("there are too many inputs")
	}

	if action != "reset" && action != "quit" {
		if len(parts) != 2 {
			return errors.New("input should be [action] [number]")
		}

		num, err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return fmt.Errorf("invalid number '%s'", parts[1])
		}
	}

	switch action {
	case "+":
		c.add(num)
	case "-":
		c.subtract(num)
	case "*":
		c.multiply(num)
	case "/":
		err = c.divide(num)
		if err != nil {
			return err
		}
	case "reset":
		c.reset()
	case "quit":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		return errors.New("action should be +, -, *, /, reset, or quit")
	}

	return nil
}
