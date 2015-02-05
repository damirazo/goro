package main

import "fmt"
import "net/http"
import "goro"
import "example"

const SERVER_PORT = "8080"

// Список маршрутов
var routes = []goro.Route{
	goro.Route{Url: `^/$`, Handler: example.Index},
	goro.Route{Url: `^/info`, Handler: example.IpAddr},
	goro.Route{Url: `^/foo`, Handler: example.FromConfig},
}

// Список параметров
var params = []goro.Config{
	goro.Config{Name: "DEBUG", Value: true},
	goro.Config{Name: "HOST", Value: "127.0.0.1"},
	goro.Config{Name: "PORT", Value: SERVER_PORT},
	goro.Config{Name: "FOO", Value: "BAR"},
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
	framework.Response = w
	framework.SetRoutes(routes)
	framework.SetParams(params)
	framework.Run()
}
