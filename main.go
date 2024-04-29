package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
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

func (stack *Stack) pop() (int, error) {
    if stack.stack_size <= 0 {
        return 0, fmt.Errorf("error: stack underflow")
    }
    stack.stack_size--
    result := stack.stack[stack.stack_size]
    stack.stack = stack.stack[:stack.stack_size]
    return result, nil
}

func (stack *Stack) peekAll() string {
    var stackContents strings.Builder
    for i := range stack.stack {
        stackContents.WriteString(strconv.Itoa(stack.stack[i]) + " ")
    }
    return stackContents.String()
}

var globalStack Stack
var commandHistory []string
var currentHistoryIndex int

func executeForth(input string, output *widget.Label) {
    words := strings.Fields(input)
    var result string

    commandHistory = append(commandHistory, input)
    currentHistoryIndex = len(commandHistory) - 1

    for _, word := range words {
        if num, err := strconv.Atoi(word); err == nil {
            globalStack.push(num)
        } else {
            if globalStack.stack_size < 2 {
                output.SetText(output.Text + "\nError: not enough elements for operation")
                return
            }

            b, err := globalStack.pop()
            if err != nil {
                output.SetText(output.Text + "\n" + err.Error())
                return
            }
            a, err := globalStack.pop()
            if err != nil {
                output.SetText(output.Text + "\n" + err.Error())
                return
            }
            switch word {
            case "+":
                globalStack.push(a + b)
            case "-":
                globalStack.push(a - b)
            case "*":
                globalStack.push(a * b)
            case "/":
                if b == 0 {
                    output.SetText(output.Text + "\nError: division by zero")
                    return
                }
                globalStack.push(a / b)
            default:
                output.SetText(output.Text + "\nUnknown command: " + word)
                return
            }
        }
    }

    result = "Stack: " + globalStack.peekAll()
    output.SetText(output.Text + "\n" + result)
}

type historyEntry struct {
    widget.Entry
}

func newHistoryEntry() *historyEntry {
    entry := &historyEntry{}
    entry.ExtendBaseWidget(entry)
    return entry
}

func (e *historyEntry) TypedKey(key *fyne.KeyEvent) {
    switch key.Name {
    case fyne.KeyUp:
        if currentHistoryIndex > 0 {
            currentHistoryIndex--
            e.SetText(commandHistory[currentHistoryIndex])
            e.CursorRow = 0
        }
    case fyne.KeyDown:
        if currentHistoryIndex < len(commandHistory)-1 {
            currentHistoryIndex++
            e.SetText(commandHistory[currentHistoryIndex])
            e.CursorRow = 0
        }
    default:
        e.Entry.TypedKey(key)
    }
}

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Forth REPL")

    input := newHistoryEntry()
    input.SetPlaceHolder("Enter Forth commands here...")

    output := widget.NewLabel("")
    output.Wrapping = fyne.TextWrapWord

    scrollContainer := container.NewVScroll(output)
    scrollContainer.SetMinSize(fyne.NewSize(780, 400))

    submitButton := widget.NewButton("Execute", func() {
        executeForth(input.Text, output)
        input.SetText("")
    })

    clearButton := widget.NewButton("Clear", func() {
        output.SetText("")
    })

    buttonContainer := container.New(layout.NewGridLayout(2), submitButton, clearButton)

    content := container.NewVBox(
        input,
        buttonContainer,
        scrollContainer,
    )

    myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(800, 600))
    myWindow.ShowAndRun()
}
