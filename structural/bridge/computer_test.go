package bridge

import "testing"

func TestMac(t *testing.T) {
	hpPrinter := &Hp{}
	macCompter := &Mac{}

	macCompter.SetPrinter(hpPrinter)
	macCompter.Print()
}
