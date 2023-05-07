package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello, World!")

	hello := widget.NewLabel("Hello, World!")
	content := container.NewVBox(hello, widget.NewButton("Quit", func() {
		myApp.Quit()
	}))
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
