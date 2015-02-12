package goro

import "regexp"
import "net/http"

// Базовая структура фреймворка
type Goro struct {
	// Входящий запрос
	Request *http.Request
	// Объект для записи с исходящие данные
	ResponseWriter *http.ResponseWriter
	// Объект исходящего запроса
	Response *Response
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
func (self *Goro) SetParams(items []ConfigItem) {
	registry := ConfigRegistry{Items: items}
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
	self.run()
}

// Поиск подходящего под URL маршрута и запуск зарегистрированного обработчика
func (self *Goro) run() {
	var handler func(*Goro)
	findedRoute := false
	routes := self.Routes
	queryString := self.Request.URL.Path

	self.Response = &Response{}

	// Производим поиск подходящего маршрута среди всего списка зарегистрированных
	for _, route := range routes {
		if matched, _ := regexp.MatchString(route.Url, queryString); matched == true {
			handler = route.Handler
			findedRoute = true
		}
	}

	// Если подходящий маршрут отсутствует, то выводим сообщение об ошибке
	if !findedRoute {
		handler = func(f *Goro) {
			//http.Error(*f.ResponseWriter, "404 page not found", http.StatusNotFound)
		}
	}

	// Запускаем обработчик
	handler(self)
	// Выбрасываем накопленную информацию клиенту
	self.Response.Flush(*self.ResponseWriter)
}

// Вывод в буфер текстовой строки
func (self *Goro) Write(s string) {
	self.Response.Write(s)
}

// Вывод в буфер текстовой строки с переносом на новую строку
func (self *Goro) WriteLine(s string) {
	self.Response.WriteLine(s)
}

// Перенаправление на указанную страницу
func (self *Goro) Redirect(url string) {
	http.Redirect(*self.ResponseWriter, self.Request, url, http.StatusOK)
}
