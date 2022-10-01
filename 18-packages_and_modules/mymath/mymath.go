package mymath

func Sum[T int | float64](a, b T) T {
	return a + b
}

// not exported
func div[T int | float64](a, b T) T {
	return a / b
}

type OperationSymbols struct {
	Sum  string
	sub  string // not visible to other packages
	Div  string
	mult string // not visible to other packages
}

func NewOperationSymbols() *OperationSymbols {
	return &OperationSymbols{
		Sum:  "+",
		sub:  "0",
		Div:  "/",
		mult: "*",
	}
}
