package goro

import "io"
import "regexp"
import "net/http"

// Базовая структура фреймворка
type Goro struct {
	// Входящий запрос
	Request *http.Request
	// Исходящий запрос
	Response *http.ResponseWriter
	// Список маршрутов
	Routes []Route
	// Список параметров
	Config ConfigRegistry
	// Объект сессии
	Session *Session
}

// Установка списка маршрутов
func (self *Goro) SetRoutes(routes []Route) {
	self.Routes = routes
}

// Установка параметров
func (self *Goro) SetParams(params []ConfigItem) {
	registry := ConfigRegistry{Data: params}
	self.Config = registry
}

// Инициализация сессий
func (self *Goro) LoadSession() {
	var uuid string
	sessionName := self.Config.Get("SESSION_NAME").(string)
	sessionCookie, err := self.Request.Cookie(sessionName)

	// Кукисы с нужным значением не нашли, возможно сессия уже устарела или открывается впервые
	if err != nil {
		uuid = GenerateUUID()
	} else {
		uuid = sessionCookie.Value
	}

	session := &Session{Id: uuid, Framework: self}

	self.Session = session
	self.Session.Load()
}

// Сохранение сессий
func (self *Goro) DumpSession() {
	self.Session.Dump()
}

// Инициализация фреймворка
func (self *Goro) Run() {
	self.findRoute()
}

// Поиск подходящего под URL маршрута и запуск зарегистрированного обработчика
func (self *Goro) findRoute() {
	routes := self.Routes
	queryString := self.Request.URL.Path

	for _, route := range routes {
		if matched, _ := regexp.MatchString(route.Url, queryString); matched == true {
			route.Handler(self)
			return
		}
	}
}

// Вывод в буфер текстовой строки
func (self *Goro) Write(s string) {
	io.WriteString(*self.Response, s)
}

// Вывод в буфер текстовой строки с переносом на новую строку
func (self *Goro) WriteLine(s string) {
	self.Write(s + "\n")
}
