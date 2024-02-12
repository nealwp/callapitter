package ui

import (
	"github.com/rivo/tview"
)

type RequestHeader struct {
	Key   string
	Value string
}

type HeadersTable struct {
	view *tview.Flex
}

func NewHeadersTable() *HeadersTable {
	view := tview.NewFlex()
	view.SetBackgroundColor(DEFAULT)
	view.SetTitle("Headers")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetBorder(true)
	return &HeadersTable{view: view}
}

func (h *HeadersTable) GetPrimitive() tview.Primitive {
	return h.view
}

func (h *HeadersTable) DisplayHeaders(headers []RequestHeader) {
	h.view.Clear()
	for _, hdr := range headers {

		key := tview.NewInputField()
		key.SetFieldBackgroundColor(DEFAULT)
		key.SetBackgroundColor(DEFAULT)
		key.SetText(hdr.Key)

		value := tview.NewInputField()
		value.SetFieldBackgroundColor(DEFAULT)
		value.SetBackgroundColor(DEFAULT)
		value.SetText(hdr.Value)

		row := tview.NewFlex()
		row.SetDirection(tview.FlexColumn)
		row.SetBackgroundColor(DEFAULT)
		row.AddItem(key, 20, 1, false)
		row.AddItem(value, 20, 1, false)
		h.view.AddItem(row, 0, 1, false)
	}
}
