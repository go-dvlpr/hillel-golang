package main

import (
	"github.com/sirupsen/logrus"
	"hillel_new/server"
)

func main() {
	err := server.Start()
	if err != nil {
		logrus.Fatal("failed to start server", err)
	}
}
