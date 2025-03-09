package trainings

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"

	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrConvToInt           = errors.New("converting into int error")
	ErrWrongLenthSlice     = errors.New("wrong length of slice error")
	ErrParseDate           = errors.New("parsing date error")
	ErrWrongDuration       = errors.New("duration can't lower than or equal 0")
	ErrUnknownTrainingType = errors.New("unknown training type")
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return ErrWrongLenthSlice
	}
	stepNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrConvToInt
	}
	t.Steps = stepNumber

	training := map[string]string{
		"Бег":    "Бег",
		"Ходьба": "Ходьба",
	}

	a, ok := training[parts[1]]
	if !ok {
		return ErrParseDate
	}
	t.TrainingType = a

	tm, err := time.ParseDuration(parts[2])
	if err != nil {
		return ErrParseDate
	}

	t.Duration = tm

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	training := []string{
		"Бег",
		"Ходьба",
	}

	if t.Duration <= 0 {
		return "", ErrWrongDuration
	}
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	if t.TrainingType == training[0] {
		energyAmount, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, energyAmount)

		return str, nil

	} else if t.TrainingType == training[1] {
		energyAmount, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		if err != nil {
			return "", err
		}

		str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, energyAmount)

		return str, nil
	} else {
		return "неизвестный тип тренировки", ErrUnknownTrainingType
	}
}
