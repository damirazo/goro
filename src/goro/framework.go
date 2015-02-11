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
	self.runRoute()
}

// Поиск подходящего под URL маршрута и запуск зарегистрированного обработчика
func (self *Goro) runRoute() {
    var handler func(*Goro)
    findedRoute := false
	routes := self.Routes
	queryString := self.Request.URL.Path

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
            http.Error(*f.Response, "404 page not found", http.StatusNotFound)
        }
    }
    
    // Запускаем обработчик
    handler(self)
}

// Вывод в буфер текстовой строки
func (self *Goro) Write(s string) {
	io.WriteString(*self.Response, s)
}

// Вывод в буфер текстовой строки с переносом на новую строку
func (self *Goro) WriteLine(s string) {
	self.Write(s + "\n")
}

// Перенаправление на указанную страницу
func (self *Goro) Redirect(url string) {
    http.Redirect(*self.Response, self.Request, url, http.StatusOK)
}
