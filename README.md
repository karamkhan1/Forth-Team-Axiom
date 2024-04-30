# Simple Forth Interpreter

This is my beta repository for my Forth interpreter written in Go. This interpreter simulates a portion of the Forth programming language.

## Usage

To use this interpreter, compile the Go source files and execute the resulting binary. The following commands illustrate how to interact with the interpreter:

- Numeric inputs (e.g., `1`, `5`) push values onto the stack.
- Arithmetic commands (e.g., `+`, `-`, `*`, `/`) perform operations on the top elements of the stack.
- Stack manipulation commands (e.g., `DUP`, `SWAP`) modify the state of the stack.

## Example Run

<img width="912" alt="Screenshot 2024-04-30 at 1 03 21 PM" src="https://github.com/karamkhan1/Forth-Team-Axiom/assets/79159011/4849495f-5a3f-46c5-9b65-75e5573ccece">


## Proposed Class Diagram
<img width="1443" alt="Object_Diagram" src="https://github.com/karamkhan1/Forth-Team-Axiom/assets/79159011/1ff2e2c3-6fad-4e92-8846-aae3fe38cd34">


## Commands

The interpreter supports the following commands for performing arithmetic operations, and manipulating the stack:

- `+`: Adds the top two elements of the stack.
- `-`: Subtracts the top element from the next top element.
- `*`: Multiplies the top two elements of the stack.
- `/`: Divides the second top element by the top element.
- `DUP`: Duplicates the top element on the stack.
- `DROP`: Removes the top element from the stack.
- `SWAP`: Swaps the top two elements of the stack.
- `OVER`: Copies the second element to the top of the stack.
- `ROT`: Rotates the third element to the top of the stack.
- `2DUP`: Duplicates the top two elements of the stack.
- `2DROP`: Removes the top two elements of the stack.
- `2SWAP`: Swaps the top two pairs of elements in the stack.
- `2OVER`: Copies the second pair of elements to the top of the stack.
- `.`: Prints the top element of the stack.
- `.S`: Prints the entire stack with the top element indicated.
- `VAR`: Introduces variable definition and assignment.

## Future Commands and Features

- Planned additions in upcoming versions include more complex data handling capabilities like arrays and structures, and enhanced control flow constructs.

For inquiries or further information, please contact me at: kkhan6@luc.edu.
