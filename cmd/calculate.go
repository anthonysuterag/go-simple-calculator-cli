/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	firstNumber  int
	secondNumber int
	operation    string
)

// calculateCmd represents the calculate command
var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "This command will perform arithmetic operations.",
	Long: `This CLI tool will calculate the two numbers that were
given as input using the mathematical operation that is passed to
the flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch operation {
		case "+":
			r := firstNumber + secondNumber
			fmt.Printf("%d + %d = %d\n", firstNumber, secondNumber, r)
		case "-":
			r := firstNumber - secondNumber
			fmt.Printf("%d - %d = %d\n", firstNumber, secondNumber, r)
		case "*":
			r := firstNumber * secondNumber
			fmt.Printf("%d * %d = %d\n", firstNumber, secondNumber, r)
		case "/":
			r := firstNumber / secondNumber
			fmt.Printf("%d / %d = %d\n", firstNumber, secondNumber, r)
		default:
			fmt.Println("Unexpected input, please provide expected operation.")
		}
	},
}

func init() {
	rootCmd.AddCommand(calculateCmd)

	calculateCmd.Flags().IntVar(&firstNumber, "input-1", 0, "Input number one.")
	calculateCmd.Flags().IntVar(&secondNumber, "input-2", 0, "Input number two.")

	calculateCmd.Flags().StringVarP(&operation, "operation", "o", "", `Arithmetic operation to be used to calculate two numbers.
	Supported input is only one of these: '+', '-', '*', '/'.`)
}
