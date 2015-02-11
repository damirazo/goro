package goro

import "os"
import "path"
import "io/ioutil"
import "encoding/json"
import "fmt"
import "errors"

var SessionItemNotFound = errors.New("Объект сессии с указанным именем не обнаружен!")

// Базовый объект для хранения информации о сессии
type Session struct {
	Id        string
	Framework *Goro
	Storage   *SessionStorage
}

// Загрузка данных из файла в объект сессии
func (self *Session) Load() {
	if self.sessionFileExist() {
		fileData, err := ioutil.ReadFile(self.sessionFilePath())
		fmt.Println(string(fileData))
		HandleError(err)
		err = json.Unmarshal(fileData, &self.Storage)
		HandleError(err)
	} else {
		self.Storage = &SessionStorage{}
	}
}

// Добавление значения в сессию
func (self *Session) Add(name string, value string) {
	self.Storage.Items = append(self.Storage.Items, &SessionItem{Name: name, Value: value})
}

// Получение значения сессии
func (self *Session) Get(name string) (string, error) {
	for _, item := range self.Storage.Items {
		if item.Name == name {
			return item.Value, nil
		}
	}
	return "", SessionItemNotFound
}

// Получение значения сессии, либо значения по умолчанию
func (self *Session) GetDefault(name string, def string) string {
	value, err := self.Get(name)
	if err == SessionItemNotFound {
		return def
	}
	return value
}

// Сохранение сессии
func (self *Session) Dump() {
	file, err := os.Create(self.sessionFilePath())
	HandleError(err)
	defer file.Close()

	data, err := json.Marshal(&self.Storage)
	fmt.Println(string(data))
	HandleError(err)
	file.Write(data)
}

// Проверка на существование файла с сессией
func (self *Session) sessionFileExist() bool {
	if _, err := os.Stat(self.sessionFilePath()); !os.IsNotExist(err) {
		return true
	}
	return false
}

// Возвращает полный путь до файла с сессией
func (self *Session) sessionFilePath() string {
	return path.Join(self.Framework.Params.Find("SESSION_PATH").(string), self.Id)
}

// Хранилище для данных сессии
type SessionStorage struct {
	Items []*SessionItem
}

// Элемент данных сессии
type SessionItem struct {
	Name  string `json: "name"`
	Value string `json: "value"`
}
