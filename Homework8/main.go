package main

import (
	"Hillel/Homework8/server"
	"github.com/sirupsen/logrus"
)

func main() {
	err := server.Start()
	if err != nil {
		logrus.Fatal("failed to start server", err)
	}
}
