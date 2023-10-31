package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/request", func(c echo.Context) error {
		get, err := http.Get("http://localhost:2000/backend")
		if err != nil {
			return err
		}
		log.Println("backend response")
		response, err := io.ReadAll(get.Body)
		log.Printf("%s", response)
		return c.String(http.StatusOK, "OK")
	})

	client := http.Client{}
	e.GET("/context_request", func(c echo.Context) error {
		req, _ := http.NewRequestWithContext(c.Request().Context(), "GET", "http://localhost:2000/backend", nil)
		log.Println("performing context request...")
		resp, err := client.Do(req)
		if err != nil {
			log.Println("error:", err)
			return err
		}

		log.Println("context backend response")
		response, err := io.ReadAll(resp.Body)
		log.Printf("%s", response)
		return c.String(http.StatusOK, "OK")
	})

	http.ListenAndServe(":1000", e)
}
