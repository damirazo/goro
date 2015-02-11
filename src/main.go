package main

import "fmt"
import "net/http"
import "goro"
import "example"
import "os"

const SERVER_PORT = "8080"

// Список маршрутов
var routes = []goro.Route{
	goro.Route{Url: `^/$`, Handler: example.Index},
	goro.Route{Url: `^/info`, Handler: example.IpAddr},
	goro.Route{Url: `^/foo`, Handler: example.FromConfig},
	goro.Route{Url: `^/session/set`, Handler: example.SessionSet},
	goro.Route{Url: `^/session/get`, Handler: example.SessionGet},
}

// Список параметров
var params = []goro.configItem{
	goro.configItem{Name: "DEBUG", Value: true},
	goro.configItem{Name: "SESSION_PATH", Value: "cache"},
	goro.configItem{Name: "SESSION_NAME", Value: "goro_session"},
	goro.configItem{Name: "HOST", Value: "127.0.0.1"},
	goro.configItem{Name: "PORT", Value: SERVER_PORT},
	goro.configItem{Name: "FOO", Value: "BAR"},
}

// Точка входа
func main() {
	http.HandleFunc("/", frameworkHandler)
	fmt.Println(http.ListenAndServe(":"+SERVER_PORT, nil))
}

// Базовый обработчик фреймворка
func frameworkHandler(w http.ResponseWriter, r *http.Request) {
	framework := &goro.Goro{}
	framework.Request = r
	framework.Response = &w
	// Загрузка маршрутов
	framework.SetRoutes(routes)
	// Загрузка параметров
	framework.SetParams(params)
	// Загрузка сессий
	framework.LoadSession()
	defer framework.DumpSession()
	// Инициализация фреймворка
	framework.Run()
	// Перехват возможных ошибок и выведение в лог
	defer func() {
		// TODO: Предусмотреть стек восстановительных задач
		e := recover()
		if e == nil {
			return
		}
		fmt.Println(e)
		os.Exit(3)
	}()
}
