package main

import (
	"fmt"

	cmd "github.com/pedrolopesme/citta-server/cmd"
)

func main() {
	err := cmd.NewBattleshipServer().Run()
	if err != nil {
		fmt.Println(err)
	}
}
