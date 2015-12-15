package tpl

import (
	"bytes"
	"tpl/layout"
)

func Index() string {
	var _buffer bytes.Buffer

	title := func() string {
		var _buffer bytes.Buffer
		"index page"
		return _buffer.String()
	}

	body := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString("<div id=\"bosyosos\">\n    </div>")

		return _buffer.String()
	}

	return layout.Base(_buffer.String(), "", "", title(), "", "")
}
