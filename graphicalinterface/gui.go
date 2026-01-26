package graphicalinterface

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Gui struct {
	app    fyne.App
	window fyne.Window
	output *widget.Entry
}

func NewGui() *Gui {
	app := app.New()

	window := app.NewWindow("mobilebg2")
	window.Resize(fyne.NewSize(1_100, 900))

	// hello := widget.NewLabel("Hello, world!")

	multilineEntry := widget.NewMultiLineEntry()
	// multilineEntry.Disable() // read-only but the text becomes unreadable
	// multilineEntry.Wrapping = fyne.TextWrapWord // wrap long lines

	container := container.NewBorder(
		nil,            // hello,          // top
		nil,            // bottom
		nil,            // left
		nil,            // right
		multilineEntry, // center (expanding)
	)

	window.SetContent(container)

	return &Gui{
		app:    app,
		window: window,
		output: multilineEntry,
	}
}

func (self *Gui) ShowAndRun() {
	self.window.ShowAndRun()
}

func (self *Gui) Print(message string) {
	fyne.Do(
		func() {
			self.output.SetText(self.output.Text + message)
		},
	)

}

func (self *Gui) InterceptStdout() {
	readOnly, writeOnly, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	os.Stdout = writeOnly

	// TODO: ideally we would check if this is already started
	go func() {
		buf := make([]byte, 1024)

		// TODO: this should not be infinite
		for {
			n, err := readOnly.Read(buf)
			if err != nil {
				panic(err)
			}

			data := string(buf[:n])
			self.Print(data)
		}
	}()
}
