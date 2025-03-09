package personaldata

import (
	"fmt"
)

// Ниже создайте структуру Personal
// Personal holds a user's personal data such as name, weight, and height.
type Personal struct{
	Name string
	Weight float64
	Height float64
}

// Ниже создайте метод Print()
// Print outputs the personal data to the console using formatted output.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}

