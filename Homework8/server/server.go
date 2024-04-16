package server

import (
	"Hillel/Homework8/config"
	"Hillel/Homework8/logger"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start() error {
	cfg := config.NewConfiguration()
	log := logger.NewLogger(cfg)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		param1 := c.QueryParam("param1")
		param2 := c.QueryParam("param2")
		log.WithFields(logrus.Fields{
			"param1": param1,
			"param2": param2,
		}).Info("GET request received")
		return c.String(http.StatusOK, "GET request received")
	})

	e.POST("/", func(c echo.Context) error {
		var requestBody RequestBody
		if err := c.Bind(&requestBody); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}
		log.WithFields(logrus.Fields{
			"bodyParam1": requestBody.BodyParam1,
			"bodyParam2": requestBody.BodyParam2,
		}).Info("POST request received")
		return c.String(http.StatusOK, "POST request received")
	})

	e.PUT("/", func(c echo.Context) error {
		var requestBody RequestBody
		if err := c.Bind(&requestBody); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}
		log.WithFields(logrus.Fields{
			"bodyParam1": requestBody.BodyParam1,
			"bodyParam2": requestBody.BodyParam2,
		}).Info("PUT request received")
		return c.String(http.StatusOK, "PUT request received")
	})

	e.DELETE("/", func(c echo.Context) error {
		var requestBody RequestBody
		if err := c.Bind(&requestBody); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}
		log.WithFields(logrus.Fields{
			"bodyParam1": requestBody.BodyParam1,
			"bodyParam2": requestBody.BodyParam2,
		}).Info("DELETE request received")
		return c.String(http.StatusOK, "DELETE request received")
	})

	log.Info("prepared to listen :1888 port")
	err := e.Start(":1888")
	if err != nil {
		log.WithError(err).Error("failed to start listening :1888 port")
		return err
	}
	return nil
}

type RequestBody struct {
	BodyParam1 string `json:"bodyParam1"`
	BodyParam2 string `json:"bodyParam2"`
}
