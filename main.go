package main

import (
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "strconv"
    "strings"
)

// here i'm defining my only object Stack to represent a stack data structure
type Stack struct {
    stack      []int // this line declares a slice to hold the stack elements
    stack_size int   // this line declares an integer to track the size of the stack
}

// push - new element is added to the stack
func (stack *Stack) push(value int) {
    stack.stack = append(stack.stack, value) // here I'm appending the value to the slice
    stack.stack_size++                        // this increments the stack size
}

// pop - top element is removed from the stack
func (stack *Stack) pop() (int, error) {
    if stack.stack_size <= 0 { // this makes sure the stack is not empty
        return 0, fmt.Errorf("error: stack underflow") // returns an error if the stack is empty
    }
    stack.stack_size--            // this decrements the stack size
    result := stack.stack[stack.stack_size] // this retrieves the top element from the stack
    stack.stack = stack.stack[:stack.stack_size] // this removes the top element from the stack
    return result, nil // returns the removed element
}

// peekAll - gets a string representation of all elements in the stack
func (stack *Stack) peekAll() string {
    var stackContents strings.Builder // here i'm declaring a strings.Builder to efficiently build the string
    for i := range stack.stack { // iterating over the stack elements
        stackContents.WriteString(strconv.Itoa(stack.stack[i]) + " ") // appending each element to the string
    }
    return stackContents.String() // returning the string representation of the stack
}

var globalStack Stack // declaring a global stack variable to be used across the program
var commandHistory []string // declaring a array to store command history
var currentHistoryIndex int // declaring a variable to track the current index in command history

// executeForth - where Forth commands are executed
func executeForth(input string, stackOutput *widget.Label, commandOutput *widget.Label) {
    words := strings.Fields(input) // splitting the input into words
    var commandValid bool = true // flag to track if the command is valid

    commandHistory = append(commandHistory, input) // adding the input to command history
    currentHistoryIndex = len(commandHistory) - 1 // updating the current history index

    for _, word := range words { // iterating over each word in the input
        if num, err := strconv.Atoi(word); err == nil { // checking if the word is a number
            globalStack.push(num) // pushing the number to the stack
        } else { // if the word is not a number, it's an operation
            switch word {
            case "+", "-", "*", "/": // if the word is an arithmetic operation
                if globalStack.stack_size < 2 { // are there enough elements in the stack?
                    stackOutput.SetText("Error: not enough elements for operation")
                    commandValid = false
                    break
                }
                b, _ := globalStack.pop() // popping the top two elements from the stack
                a, _ := globalStack.pop()
                switch word {
                case "+": // performing addition
                    globalStack.push(a + b)
                case "-": // performing subtraction
                    globalStack.push(a - b)
                case "*": // performing multiplication
                    globalStack.push(a * b)
                case "/": // performing division
                    if b == 0 {
                        stackOutput.SetText("Error: division by zero")
                        commandValid = false
                    } else {
                        globalStack.push(a / b)
                    }
                }
        case "DUP":
            // top element of the stack is duplicated
            if globalStack.stack_size < 1 { // this makes sure there's at least one element in the stack
                stackOutput.SetText("Error: not enough elements to duplicate") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            top := globalStack.stack[globalStack.stack_size-1] // retrieving the top element of the stack
            globalStack.push(top) // pushing a copy of the top element onto the stack
        
        case "DROP":
            // top element of the stack is removed
            if globalStack.stack_size < 1 { // this makes sure there's at least one element in the stack
                stackOutput.SetText("Error: not enough elements to drop") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            globalStack.pop() // removing the top element from the stack
        
        case "SWAP":
            // top two elements of the stack are swapped
            if globalStack.stack_size < 2 { // this makes sure there are at least two elements in the stack
                stackOutput.SetText("Error: not enough elements to swap") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            a, _ := globalStack.pop() // popping the top two elements from the stack
            b, _ := globalStack.pop()
            globalStack.push(a) // pushing them back in reverse order
            globalStack.push(b)
        
        case "OVER":
            // second element from the top of the stack is duplicated
            if globalStack.stack_size < 2 { // this makes sure there are at least two elements in the stack
                stackOutput.SetText("Error: not enough elements for operation") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            a := globalStack.stack[globalStack.stack_size-2] // retrieving the second element from the top of the stack
            globalStack.push(a) // pushing a copy of it onto the stack
        
        case "ROT":
            // third element from the top of the stack is rotated to the top
            if globalStack.stack_size < 3 { // this makes sure there are at least three elements in the stack
                stackOutput.SetText("Error: not enough elements for rotation") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            c, _ := globalStack.pop() // popping the top three elements from the stack
            b, _ := globalStack.pop()
            a, _ := globalStack.pop()
            globalStack.push(b) // pushing them back in a rotated order
            globalStack.push(c)
            globalStack.push(a)
        
        case "2DUP":
            // top two elements of the stack are duplicated
            if globalStack.stack_size < 2 { // this makes sure there are at least two elements in the stack
                stackOutput.SetText("Error: not enough elements to duplicate") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            b := globalStack.stack[globalStack.stack_size-1] // retrieving the top two elements from the stack
            a := globalStack.stack[globalStack.stack_size-2]
            globalStack.push(a) // pushing copies of them back onto the stack
            globalStack.push(b)
        
        case "2DROP":
            // top two elements of the stack are removed
            if globalStack.stack_size < 2 { // this makes sure there are at least two elements in the stack
                stackOutput.SetText("Error: not enough elements to drop") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            globalStack.pop() // popping the top two elements from the stack
            globalStack.pop()
        
        case "2SWAP":
            // top four elements of the stack are swapped
            if globalStack.stack_size < 4 { // this makes sure there are at least four elements in the stack
                stackOutput.SetText("Error: not enough elements to swap") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            d, _ := globalStack.pop() // popping the top four elements from the stack
            c, _ := globalStack.pop()
            b, _ := globalStack.pop()
            a, _ := globalStack.pop()
            globalStack.push(c) // pushing them back in a swapped order
            globalStack.push(d)
            globalStack.push(a)
            globalStack.push(b)
        
        case "2OVER":
            // third and fourth elements from the top of the stack are duplicated
            if globalStack.stack_size < 4 { // this makes sure there are at least four elements in the stack
                stackOutput.SetText("Error: not enough elements for operation") // setting error message if not enough elements
                commandValid = false // marking command as invalid
                break
            }
            b := globalStack.stack[globalStack.stack_size-2] // retrieving the third and fourth elements from the top of the stack
            a := globalStack.stack[globalStack.stack_size-3]
            globalStack.push(a) // pushing copies of them back onto the stack
            globalStack.push(b)
        
        default:
            // if not all that sits above its an unknown command
            stackOutput.SetText("Unknown command: " + word) // setting error message for unknown command
            commandValid = false // marking command as invalid
            break
        }
        

            if !commandValid { // exiting loop if command is invalid
                break
            }
        }
    }

    // updating the command output label with the result of the command
    if commandValid {
        commandOutput.SetText(commandOutput.Text + "\n" + input + " ok")
    } else {
        commandOutput.SetText(commandOutput.Text + "\n" + input + " ?")
    }

    stackOutput.SetText(globalStack.peekAll() + "<- Top") // updating the stack output label
}

// here i defined a custom entry widget to handle command history navigation and execution - bonus points
type historyEntry struct {
    widget.Entry // embedding the standard Entry widget
    stackOutput   *widget.Label // storing a reference to the stack output label
    commandOutput *widget.Label // storing a reference to the command output label
}

// this is the constructor function for the custom history entry widget
func newHistoryEntry(stackOutput *widget.Label, commandOutput *widget.Label) *historyEntry {
    entry := &historyEntry{stackOutput: stackOutput, commandOutput: commandOutput}
    entry.ExtendBaseWidget(entry) // extending the base widget with custom functionality
    return entry
}

// this is the part of the program where key events are handled for the history entry widget
func (e *historyEntry) TypedKey(key *fyne.KeyEvent) {
    switch key.Name {
    case fyne.KeyUp: // navigating to previous command in history
        if currentHistoryIndex > 0 {
            currentHistoryIndex--
            e.SetText(commandHistory[currentHistoryIndex])
            e.CursorRow = 0
        }
    case fyne.KeyDown: // navigating to next command in history
        if currentHistoryIndex < len(commandHistory)-1 {
            currentHistoryIndex++
            e.SetText(commandHistory[currentHistoryIndex])
            e.CursorRow = 0
        }
    case fyne.KeyReturn, fyne.KeyEnter: // executing the command
        executeForth(e.Text, e.stackOutput, e.commandOutput)
        e.SetText("")
    default:
        e.Entry.TypedKey(key) // executing the input
    }
}

func main() {
    // initializing the fyne application and window
    myApp := app.New()
    myWindow := myApp.NewWindow("Forth REPL")

    // creating labels to display stack and command output
    stackOutput := widget.NewLabel("")
    stackOutput.Wrapping = fyne.TextWrapWord

    background := canvas.NewRectangle(theme.BackgroundColor())
    background.FillColor = theme.InputBackgroundColor()
    stackContainer := container.NewMax(background, stackOutput)

    commandOutput := widget.NewLabel("")
    commandOutput.Wrapping = fyne.TextWrapWord

    // creating a custom entry widget for input with command history support
    input := newHistoryEntry(stackOutput, commandOutput)
    input.SetPlaceHolder("Enter Forth commands here...")

    // creating scrollable containers for stack and command output
    stackScroll := container.NewVScroll(stackContainer)
    stackScroll.SetMinSize(fyne.NewSize(780, 200))

    commandScroll := container.NewVScroll(commandOutput)
    commandScroll.SetMinSize(fyne.NewSize(780, 200))

    // creating buttons for executing and clearing commands
    submitButton := widget.NewButton("Execute", func() {
        executeForth(input.Text, stackOutput, commandOutput)
        input.SetText("")
    })

    clearButton := widget.NewButton("Clear", func() {
        commandOutput.SetText("")
    })

    // arranging buttons in a grid layout
    buttonContainer := container.New(layout.NewGridLayout(2), submitButton, clearButton)

    // arranging widgets vertically
    content := container.NewVBox(
        input,
        buttonContainer,
        stackScroll,
        commandScroll,
    )

    // setting content and size for the window
    myWindow.SetContent(content)
    myWindow.Resize(fyne.NewSize(800, 600))
    myWindow.ShowAndRun() // displaying the window and starting the event loop
}
