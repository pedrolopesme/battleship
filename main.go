package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pedrolopesme/citta-server/core"
)

func main() {

	broker := core.NewServer()

	go func() {
		for {
			//time.Sleep(time.Millisecond * 500)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal("HTTP server error: ", http.ListenAndServe(":3030", broker))

}
