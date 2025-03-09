package actioninfo

import (
	"errors"
	"fmt"
)

// DataParser — интерфейс, который должны реализовать ваши структуры Training и DaySteps.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Возможные ошибки
var (
	ErrParsing       = errors.New("parsing error")
	ErrActionInfo    = errors.New("action info error")
)

// Info принимает слайс строк (dataset) и экземпляр dp, реализующий интерфейс DataParser.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Пытаемся распарсить строку
		if err := dp.Parse(data); err != nil {
			// Выводим информацию о том, что возникла ошибка парсинга
			fmt.Println(ErrParsing, err)
			continue
		}

		// Формируем информацию об активности
		info, err := dp.ActionInfo()
		if err != nil {
			// Выводим информацию о том, что возникла ошибка при формировании данных
			fmt.Println(ErrActionInfo, err)
			continue
		}

		// Выводим результат
		fmt.Println(info)
	}
}
