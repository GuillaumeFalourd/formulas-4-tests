package formula

import (
	"fmt"
)

type Formula struct {
	NumberOne int
	NumberTwo int
}

func (h Formula) Run() {
	sum := h.NumberOne + h.NumberTwo
	fmt.Println("Sum:", sum)
}
