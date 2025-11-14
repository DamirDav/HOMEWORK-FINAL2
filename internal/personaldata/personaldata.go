package personaldata

import "fmt"

type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес, кг
	Height float64 // рост, м
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
