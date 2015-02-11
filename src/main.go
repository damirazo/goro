package main

import "fmt"
import "goro"
import "net/http"
import "example"

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
var params = []goro.ConfigItem{
	goro.ConfigItem{Name: "DEBUG", Value: true},
	goro.ConfigItem{Name: "SESSION_PATH", Value: "cache"},
	goro.ConfigItem{Name: "SESSION_NAME", Value: "goro_session"},
	goro.ConfigItem{Name: "HOST", Value: "127.0.0.1"},
	goro.ConfigItem{Name: "PORT", Value: SERVER_PORT},
	goro.ConfigItem{Name: "FOO", Value: "BAR"},
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
}
