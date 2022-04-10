package main

import (
	"fynetodo/fynetodo/ui"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Todo")
	w.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})

	id := 1
	todoListForm := ui.NewTodoListForm()
	addTodoBtn := widget.NewButton("Add", func() {
		entry := widget.NewEntry()
		input := widget.NewFormItem("Todo", entry)
		dialog.ShowForm("Add todo", "Save", "Cancel", []*widget.FormItem{input}, func(b bool) {
			text := entry.Text
			if strings.Trim(text, " ") == "" {
				return
			}
			todoListForm.Add(id, text)
			id++
		}, w)
	})

	content := container.New(layout.NewVBoxLayout(), todoListForm.Container(), addTodoBtn)
	w.SetContent(content)
	w.ShowAndRun()
}
