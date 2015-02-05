package goro

// Системный маршрут
type Route struct {
	Url     string
	Handler func(*Goro)
}
