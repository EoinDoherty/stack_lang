package runtime

const (
    END_STR = iota // Null terminator for strings
	PUSH           // Pushes a number onto the stack
	POP            // Pops a value from the stack
    POP_STR        // Pops the string at the top of the stack
	ADD            // Adds first two values on the stack
	SUB            // Subtracts second stack value from the first
	DIV            // Divides first stack value by the second
	MUL            // Multiplies first two values on the stack
    PRINT_VAL      // Prints the top number on the stack
    PRINT_STR      // Prints the top string on the stack
)
