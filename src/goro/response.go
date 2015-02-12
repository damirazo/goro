package goro

import "io"
import "net/http"

// Объект ответа от сервера
type Response struct {
	StatusCode int
	Content    string
}

// Дополнение ответа данными
func (self *Response) Write(s string) {
	self.Content += s
}

// Дополнение ответа текстовой строкой
func (self *Response) WriteLine(s string) {
	self.Content += s + "\n"
}

// Отправка ответа
func (self *Response) Flush(r http.ResponseWriter) {
	io.WriteString(r, self.Content)
}
