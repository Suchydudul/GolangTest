package new

// 	var Atr int = rand.Intn(10-1)
//fmt.Println(Atr)
import (
	"fmt"
	"rand"
)

func new(int) int {
	var atr int = rand.Intn(10 - 1)
	fmt.Println("punkty atrybutów: %s", atr)
	return atr
}
