package ui

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type TodoListForm struct {
	parent *fyne.Container
	todos  map[int]*fyne.Container
}

func NewTodoListForm() *TodoListForm {
	c := container.New(layout.NewGridLayout(1))
	todos := make(map[int]*fyne.Container)
	return &TodoListForm{parent: c, todos: todos}
}

func (t *TodoListForm) Container() fyne.CanvasObject {
	return t.parent
}

func (t *TodoListForm) Add(id int, value string) {
	text := canvas.NewText(value, color.White)
	idStr := fmt.Sprintf("%d", id)
	radioBtn := widget.NewRadioGroup([]string{idStr}, func(s string) {
		id, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		todo, ok := t.todos[id]
		if ok {
			t.parent.Remove(todo)
			delete(t.todos, id)
		}
	})
	newTodo := container.New(layout.NewHBoxLayout(), radioBtn, text)
	t.parent.AddObject(newTodo)
	t.todos[id] = newTodo
}
