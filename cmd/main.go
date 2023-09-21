package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Print("Test fyne.")
	a := app.New()
	w := a.NewWindow("SongPlayer")
	w.ShowAndRun()
}
