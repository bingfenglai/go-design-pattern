package iterator_dp

import (
	"errors"
)

type Iterator interface {
	HasNext() bool
	Next() (error, interface{})
}

type HdlMenuIterator struct {
	menuItem []MenuItem
	index    int
}

func NewHdlMenuIterator(menuItem []MenuItem) *HdlMenuIterator {
	return &HdlMenuIterator{menuItem: menuItem, index: 0}
}

func (h *HdlMenuIterator) HasNext() bool {
	return h.index < len(h.menuItem)
}

func (h *HdlMenuIterator) Next() (error, interface{}) {
	if h.HasNext() {
		item := h.menuItem[h.index]
		h.index++
		return nil, item
	}
	return errors.New("数组越界"), nil
}
