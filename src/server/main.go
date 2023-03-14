package main

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sheodox/mantra/mantras"
	"github.com/sheodox/mantra/message"
)

const (
	// general starting point used for mantra timing calculations
	BASE_MINIMUM_WAIT_TIME = 4 * time.Hour
	// absolute shortest amount of time between mantras
	MINIMUM_WAIT_TIME = 15 * time.Minute
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
	// We should generally wait a while before the next mantra, but this lets mantras
	// occasionally be sent more often.
	minimum := BASE_MINIMUM_WAIT_TIME * time.Duration(getMinimumWaitTimeModifier())

	// wait some randomish amount of time before sending the next mantra so it's never expected
	// and hopefully isn't so often it loses all impact or gets ignored
	extraTime := time.Duration(rand.Intn(8))*time.Hour + time.Duration(rand.Intn(60))*time.Minute

	sleepTime := extraTime + minimum

	// make sure we're not sending a new mantra too soon, it's very unlikely
	// but still possible that we could spam a bunch of mantras if the minimum modifier
	// and extra time were both low several times in a row.
	if sleepTime < MINIMUM_WAIT_TIME {
		sleepTime = MINIMUM_WAIT_TIME
	}

	time.Sleep(sleepTime)
}

func getMinimumWaitTimeModifier() float64 {
	x := rand.Float64()
	// this is equivalent to doing a base 500 log, which gets a much steeper curve.
	// so 90% of the time we're waiting 2/3 of the minimum at least, but for small
	// x values it can dip all the way down towards 0 (under 0.002 it is negative,
	// but it won't wait a time below MINIMUM_WAIT_TIME)
	// x = 0.003 return 0.1
	// x = 0.025 return 0.4
	// x = 0.05 return 0.52
	// x = 0.1 return 0.63
	// x = 0.2 return 0.74
	// x = 0.5 return 0.88
	// x = 1.0 return 1
	return math.Log(x)/math.Log(500) + 1
}

func sendRandomMantra(m message.Messager) {
	mantra, hasMantras := mantras.GetRandomMantra()

	if hasMantras {
		m.SendMessage(mantra.Text)
	}
}
