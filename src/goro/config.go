package goro

// Регистр параметров
type ConfigRegistry struct {
	Items []ConfigItem
}

// Ищем параметр по имени
func (self *ConfigRegistry) Get(paramName string) interface{} {
	for _, config := range self.Items {
		if paramName == config.Name {
			return config.Value
		}
	}

	return nil
}

// Параметр
type ConfigItem struct {
	Name  string
	Value interface{}
}
