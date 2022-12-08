package memento_dp

import "log"

type EditMemento struct {
	content string
}

type Editor struct {
	content string
}

func (receiver *Editor) SetContent(content string) {
	receiver.content = content
}

func NewEditor(content string) *Editor {
	return &Editor{content: content}
}

func (receiver *Editor) CreateMemento() *EditMemento {
	return &EditMemento{content: receiver.content}
}

func (receiver *Editor) Rollback(memento *EditMemento) {
	receiver.content = memento.content
}

func (receiver *Editor) ShowContent() {
	log.Default().Println("content: ", receiver.content)
}

type Caretaker struct {
	memento *EditMemento
}

func (c *Caretaker) Memento() *EditMemento {
	return c.memento
}

func NewCaretaker(originator *Editor) *Caretaker {
	return &Caretaker{memento: originator.CreateMemento()}
}
