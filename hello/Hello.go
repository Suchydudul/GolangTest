package main

//tersdsa
import (
	"fmt"
	"github.com/Suchydudul/GolangTest/new"
)

func main() {

	var par string
	fmt.Println("N/Q")
	fmt.Scanln(&par)
	if par == "N" {
		fmt.Println("Kabanos ♥")
		message := new.newgame(1)
		fmt.Println(message)

	} else if par == "S" {
		fmt.Println("ser :CCCC")
	}

}
