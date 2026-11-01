package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("ошибка при парсинге: %v\n", err)
			continue
		}
		s, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка при формировании строки: %v\n", err)
		}
		fmt.Println(s)
	}
}
