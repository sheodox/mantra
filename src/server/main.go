package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sheodox/mantra/mantras"
	"github.com/sheodox/mantra/message"
)

const (
	MINIMUM_WAIT_TIME = 4 * time.Hour
)

var config Config

func main() {
	rand.Seed(int64(time.Now().UnixNano()))

	config = loadConfig()
	messager := message.NewMessager(config.DiscordWebhookUrl)

	go scheduleNextMantra(messager)

	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/mantras", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mantras.GetMantras())
	})
	e.POST("/api/mantras", func(c echo.Context) error {
		err := mantras.AddMantra(c.FormValue("text"))
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusOK)
	})
	e.GET("/api/mantras/:id", func(c echo.Context) error {
		id := c.Param("id")

		mantra, err := mantras.GetMantra(id)

		if err != nil {
			if err == mantras.ErrNotFound {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusBadRequest)
		}

		return c.JSON(http.StatusOK, mantra)
	})
	e.POST("/api/mantras/:id", func(c echo.Context) error {
		id := c.Param("id")
		err := mantras.UpdateMantra(c.FormValue("text"), id)

		if err != nil {
			if err == mantras.ErrNotFound {
				return c.NoContent(http.StatusNotFound)
			}
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusOK)
	})
	e.DELETE("/api/mantras/:id", func(c echo.Context) error {
		id := c.Param("id")
		err := mantras.DeleteMantra(id)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func scheduleNextMantra(m message.Messager) {
	sendRandomMantra(m)

	sleepUntilNextMantra()
	scheduleNextMantra(m)
}

func sleepUntilNextMantra() {
	// wait some randomish amount of time before sending the next mantra so it's never expected
	// and hopefully isn't so often it loses all impact or gets ignored
	extraTime := time.Duration(rand.Intn(8))*time.Hour + time.Duration(rand.Intn(60))*time.Minute

	sleepTime := extraTime + MINIMUM_WAIT_TIME

	time.Sleep(sleepTime)
}

func sendRandomMantra(m message.Messager) {
	mantra, hasMantras := mantras.GetRandomMantra()

	if hasMantras {
		m.SendMessage(mantra.Text)
	}
}
