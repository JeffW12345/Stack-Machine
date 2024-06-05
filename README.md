Stack Machine
=============

README below adapated from https://github.com/bjssacademy/stackmachine. 

This program implements a stack based interpreter using Golang. This is a tiny version of a programming language interpreter.

This machine is given a string, containing a sequence of commands. Each command is separated by one space. The commands operate on a stack of integers. 

**Commands to implement**

The following commands must be implemented:

An integer (from 0 to 50000): the machine pushes this integer onto the stack
"POP": the machine removes the most recently pushed number from the stack
"DUP": duplicate the last number pushed on the stack (or keep the stack empty)
"+": pop the most recent two numbers; add them together; push the result. If an overflow occurs (result higher than 50000) return an error
"-": pop the most recent two numbers; subtract the second from the first most recent. push the result. If the result is below zero, return an error
"*" pop the top two elements off the stack, multiply them together, push the result onto the stack
"CLEAR" empties the stack so it has no elements on it
"SUM" pops all elements off the stack, adds them together, and pushes the result onto the stack. SUM on an empty stack returns an error.

Any other input is invalid. the machine must stop and return an error.

After processing all the operations without errors, the machine returns the topmost value from the stack. If an error occurs at any point, the machine stops further processing an returns an error.

All numbers must lie between 0 and 50,000 inclusive. Any result that lies outside this range must stop the machine and return an error.

**Example command sequences**

The following examples will clarify what the machine does.

You must write a test for every one. All tests must pass.

Example: 99
The command "99" will:

Push the value 99 onto the stack
Result: The machine will pop the most recent value off the stack (99) and return it, without error

Example: 1 2 +
The command sequence "1 2 +" will:

Push 1 onto the stack
Push 2 onto the stack
pop the numbers off the stack and add them, forming a result of 3
push this 3 onto the stack
The stack will contain the single value of 3.

Result: The machine will return this value of 3 with no error.

Example: 3 DUP +
The command sequence "3 DUP +" will:

Push 3 onto the stack
Duplicate the most recent value on the stack, making the stack contain [3, 3]
Plus will take both number off the stack, add them together giving 6 and
push this result (6) onto the stack
Result: The machine will return 6 and no error

Example: DUP 99
The command sequence "DUP 99" will:

duplicate the topmost number on the stack. The stack is empty, so it will leave the stack being empty
Push 99 onto the stack
Result: The machine will return 99 and no error

Example: 3 DUP * 1 +
The command sequence "3 DUP * 1 +" will:

Push 3 onto the stack
Duplicate the 3
Pop both values off the stack and multiply them, giving a result of 9
Push this result (9) onto the stack
Push 1 onto the stack
Add the most recent two numbers on the stack giving 10
Push this result (10) onto the stack
Result: The machine will return 10 and no error

Example: 50000 1 +
the command sequence "50000 1 +" will:

Push 50000 onto the stack
Push 1 onto the stack
Pop the two numbers off the stack and add them
The result is greater than 50,000 which is an overflow
Result: The machine will stop and report an error

Example 50001
The command "50001" will:

Stop the machine due to out of range value
Result: The machine will stop and report an error

Example: 1 2 3 4 5 SUM
The command sequence "1 2 3 4 5 SUM" will:

Push the numbers 1, 2, 3, 4 and 5 onto the stack
Pop all numbers on the stack and add them up
Push the result (15) onto the stack
Result: The machine will retun 15 and no error

Example: XXX-INVALID 1 2 +
The command "XXX-INVALID" will:

Stop the machine as XXX-INVALID is not a valid input
Result: The machine will stop and report an error

Example: 2 5 -
The commands "2 5 -" will:

Push 2 onto the stack
Push 5 onto the stack
Pop off the two numbers and subtract
The result 3 will be pushed on the stack
Result: The machine will return 3 and no error

INSTRUCTIONS
============

These instructions assume that you have Golang installed on your device. 

Download the repo. Within a console window, type:

go run .

to run the code. 

You can adapt the code in the main method in stack_machine_main.go to get the outputs from various stack instruction input strings.

To run the tests, enter:

go test
