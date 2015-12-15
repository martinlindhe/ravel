package layout

import (
	"bytes"
)

func Base(body string, sidebar string, footer string, title string, css string, js string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\" />\n    <title>")
	_buffer.WriteString((title))
	_buffer.WriteString("</title>\n</head>\n<body>\n    <div class=\"container\">")
	_buffer.WriteString((body))
	_buffer.WriteString("</div>\n    <div class=\"sidebar\">")
	_buffer.WriteString((sidebar))
	_buffer.WriteString("</div>\n    <div class=\"footer\">")
	_buffer.WriteString((footer))
	_buffer.WriteString("</div>\n    ")
	_buffer.WriteString((js))
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer.String()
}
