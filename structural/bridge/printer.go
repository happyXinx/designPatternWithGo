package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("epson printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("hp printer")
}
