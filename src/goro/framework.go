package goro

import "io"
import "net/http"
import "fmt"
import "time"
import "regexp"

// Базовая структура фреймворка
type Goro struct {
	// Входящий запрос
	Request *http.Request
	// Исходящий запрос
	Response http.ResponseWriter
	// Список маршрутов
	Routes []Route
	// Список параметров
	Params ConfigRegistry
}

// Установка списка маршрутов
func (self *Goro) SetRoutes(routes []Route) {
	self.Routes = routes
}

// Установка параметров
func (self *Goro) SetParams(params []Config) {
	registry := ConfigRegistry{Params: params}
	self.Params = registry
}

// Инициализация фреймворка
func (self *Goro) Run() {
	self.FindRoute()
}

// Поиск подходящего под URL маршрута и запуск зарегистрированного обработчика
func (self *Goro) FindRoute() {
	routes := self.Routes
	queryString := self.Request.URL.Path
	now := time.Now().Local().Format("2006-01-02 15:04:05")
	ip := self.Request.RemoteAddr

	finded := false
	returnCode := http.StatusOK

	for _, route := range routes {
		if matched, _ := regexp.MatchString(route.Url, queryString); matched == true {
			route.Handler(self)
			finded = true

		}
	}

	// Если соответствие не найдено, то возвращаем 404 код
	if !finded {
		returnCode = http.StatusNotFound
	}

	self.Response.WriteHeader(returnCode)
	fmt.Println("[", now, "] ", ip, "|", queryString, "|", returnCode)
}

// Вывод в буфер текстовой строки
func (self *Goro) Write(s string) {
	io.WriteString(self.Response, s)
}

// Вывод в буфер текстовой строки с переносом на новую строку
func (self *Goro) WriteLine(s string) {
	self.Write(s + "\n")
}
