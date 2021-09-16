package main

import (
	"fmt"

	cmd "github.com/pedrolopesme/battleship/cmd"
)

func main() {
	err := cmd.NewBattleshipServer().Run()
	if err != nil {
		fmt.Println(err)
	}
}
