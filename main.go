package main

import (
	"github.com/bhaktiutama/banking/app"
	"github.com/bhaktiutama/banking/logger"
)

func main() {
	// log.Println("Starting out application ...")

	logger.Info("Starting the application...")
	app.Start()
}
