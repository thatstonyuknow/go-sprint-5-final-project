package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
)

// StepLength — средняя длина шага в метрах.
const (
	StepLength = 0.65
)

var (
	ErrWrongDataFormat = errors.New("wrong length of data string")
	ErrConvSteps       = errors.New("can't convert steps to int")
	ErrParseDuration   = errors.New("can't parse duration")
)

// DaySteps содержит данные о дневных прогулках.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку формата "steps,duration" (например, "678,0h50m")
// и сохраняет результат в структуру DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return ErrWrongDataFormat
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return ErrConvSteps
	}
	ds.Steps = steps

	// Парсим длительность
	dur, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return ErrParseDuration
	}
	ds.Duration = dur

	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Вычислим дистанцию (в километрах)
	distance := (float64(ds.Steps) * StepLength) / 1000.0

	// Формируем итоговую строку
	info := fmt.Sprintf(`Имя: %s
Вес: %.2f
Рост: %.2f
Шаги: %d
Длительность: %.2f ч.
Дистанция: %.2f км.
`, ds.Name, ds.Weight, ds.Height, ds.Steps, ds.Duration.Hours(), distance)

	return info, nil
}
