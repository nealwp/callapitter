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
	view.SetBackgroundColor(BG_COLOR)
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
		key.SetFieldBackgroundColor(BG_COLOR)
		key.SetBackgroundColor(BG_COLOR)
		key.SetText(hdr.Key)

		value := tview.NewInputField()
		value.SetFieldBackgroundColor(BG_COLOR)
		value.SetBackgroundColor(BG_COLOR)
		value.SetText(hdr.Value)

		row := tview.NewFlex()
		row.SetDirection(tview.FlexColumn)
		row.SetBackgroundColor(BG_COLOR)
		row.AddItem(key, 20, 1, false)
		row.AddItem(value, 20, 1, false)
		h.view.AddItem(row, 0, 1, false)
	}
}
