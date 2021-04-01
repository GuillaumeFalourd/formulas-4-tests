package formula

import (
	"fmt"
)

type Formula struct {
	NumberOne int
	NumberTwo int
}

func (h Formula) Run() {
	multiplication := h.NumberOne * h.NumberTwo
	fmt.Println("Multiplication:", multiplication)
}
