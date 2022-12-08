package test

import (
	"lbf.com/go-dp/memento_dp"
	"testing"
)

func TestMemento(t *testing.T) {
	editor := memento_dp.NewEditor("Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。")
	caretaker := memento_dp.NewCaretaker(editor)
	editor.ShowContent()
	editor.SetContent("Go 是一个开源的编程语言.")
	editor.ShowContent()
	editor.Rollback(caretaker.Memento())
	editor.ShowContent()

}
