package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		for {
			event := "Event " + time.Now().String()
			if err := json.NewEncoder(c.Response()).Encode(event); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(1 * time.Second)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
