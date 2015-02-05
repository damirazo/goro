package goro

// Регистр параметров
type ConfigRegistry struct {
	Params []Config
}

// Ищем параметр по имени
func (self *ConfigRegistry) Find(paramName string) interface{} {
	for _, config := range self.Params {
		if paramName == config.Name {
			return config.Value
		}
	}

	return nil
}

// Параметр
type Config struct {
	Name  string
	Value interface{}
}
