package parser

// Values of the AST slice
const (
	END_STR   = iota // Null terminator for strings (Does this need to be in the spec?)
	PUSH             // Pushes a number onto the stack
	POP              // Pops a value from the stack
	ADD              // Adds first two values on the stack
	SUB              // Subtracts second stack value from the first
	DIV              // Divides first stack value by the second
	MUL              // Multiplies first two values on the stack
	PRINT_VAL        // Prints the top number on the stack
	PRINT_STR        // Prints the top string on the stack
)
