package helper

import (
	"bytes"
)

func Footer() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<div>\n    footer\n</div>")

	return _buffer.String()
}
