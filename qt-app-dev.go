// +build dev

package main

import (
	"os"
	fp "path/filepath"

	"github.com/go-qamel/qamel"
	"github.com/sirupsen/logrus"
)

func runQtApp(argc int, argv []string) {
	logrus.Println("DEV MODE")

	// Create QT app
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("JSON to Go")
	app.SetWindowIcon(":/res/icon.png")

	// Create viewer
	view := qamel.NewViewer()
	view.SetSource("res/main.qml")
	view.SetResizeMode(qamel.SizeRootObjectToView)
	view.SetHeight(600)
	view.SetWidth(800)
	view.ShowMaximized()

	// Watch change in resource dir
	projectDir, err := os.Getwd()
	if err != nil {
		logrus.Fatalln("Failed to get working directory:", err)
	}

	resDir := fp.Join(projectDir, "res")
	go view.WatchResourceDir(resDir)

	// Exec app
	app.Exec()
}
