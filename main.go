package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

type Stack struct {
	stack      []int
	stack_size int
}

func (stack *Stack) push(value int) {
	stack.stack = append(stack.stack, value)
	stack.stack_size++
}

func (stack *Stack) pop() int {
	if stack.stack_size <= 0 {
		panic("error: stack underflow")
	}
	stack.stack_size--
	result := stack.stack[stack.stack_size]
	stack.stack = stack.stack[:stack.stack_size]
	return result
}

func (stack *Stack) peekAll() string {
	var stackContents strings.Builder
	for i := range stack.stack {
		stackContents.WriteString(strconv.Itoa(stack.stack[i]) + " ")
	}
	return stackContents.String()
}

var globalStack Stack

func executeForth(input string, output *widget.Label) {
	words := strings.Fields(input)
	var result string

	for _, word := range words {
		if num, err := strconv.Atoi(word); err == nil {
			globalStack.push(num)
		} else {
			if globalStack.stack_size < 2 {
				result = "Error: not enough elements for operation"
				output.SetText(output.Text + "\n" + result)
				return
			}

			b := globalStack.pop()
			a := globalStack.pop()
			switch word {
			case "+":
				globalStack.push(a + b)
			case "-":
				globalStack.push(a - b)
			case "*":
				globalStack.push(a * b)
			case "/":
				if b == 0 {
					result = "Error: division by zero"
					output.SetText(output.Text + "\n" + result)
					return
				}
				globalStack.push(a / b)
			default:
				result = "Unknown command: " + word
				output.SetText(output.Text + "\n" + result)
				return
			}
		}
	}

	result = "Stack: " + globalStack.peekAll()
	output.SetText(output.Text + "\n" + result)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Forth REPL")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Forth commands here...")

	output := widget.NewLabel("")
	output.Wrapping = fyne.TextWrapWord

	submitButton := widget.NewButton("Execute", func() {
		executeForth(input.Text, output)
		input.SetText("") // Clear the input after executing the command
	})

	content := container.NewVBox(
		input,
		submitButton,
		output,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
