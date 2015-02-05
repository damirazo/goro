package example

import "goro"

// Вывод текстовой строки
func Index(f *goro.Goro) {
	f.WriteLine("Hello, World!")
}

// Вывод информации об клиенте
func IpAddr(f *goro.Goro) {
	f.WriteLine("Ваш ip адрес: " + f.Request.RemoteAddr)
	f.WriteLine("Информация о браузере: " + f.Request.UserAgent())
}

// Отображение значения из параметров
func FromConfig(f *goro.Goro) {
	f.WriteLine(f.Params.Find("FOO").(string))
}
