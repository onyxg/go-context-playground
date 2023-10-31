package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/backend", func(c echo.Context) error {

		log.Println("backend query...")

		err := doWork(c.Request().Context())
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	http.ListenAndServe(":2000", e)
}

func doWork(ctx context.Context) error {
	funcComplete := make(chan bool)
	go func() {
		log.Println("starting backend query...")
		time.Sleep(5 * time.Second)
		funcComplete <- true
		log.Println("backend query completed...")
	}()

	for {
		select {
		case <-ctx.Done():
			// The context is over, stop processing results
			log.Println("context done, cancelling request....")
			return errors.New("context done")
		case <-funcComplete:
			// Process the results received
			return nil
		}
	}
}
