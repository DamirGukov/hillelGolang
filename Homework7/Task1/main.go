package main

import (
	"Hillel/Homework7/Task1/server"
	"github.com/sirupsen/logrus"
)

func main() {
	err := server.Start()
	if err != nil {
		logrus.Fatal("failed to start server", err)
	}
}
