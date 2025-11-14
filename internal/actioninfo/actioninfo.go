package actioninfo

import "log"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, line := range dataset {
		if err := dp.Parse(line); err != nil {
			log.Printf("ошибка парсинга: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		print(info)
	}
}
