package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(source string) error {
	parts := strings.Split(source, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 { //  положительные шаги
		return errors.New("количество шагов должно быть положительным числом")
	}
	t.Steps = steps
	t.TrainingType = parts[1]

	dur, err := time.ParseDuration(parts[2])
	if err != nil || dur <= 0 { // положительная длительность
		return errors.New("продолжительность должна быть положительной")
	}
	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	switch t.TrainingType {
	case "Бег", "бег":
		cal, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			spentenergy.Distance(t.Steps, t.Height),
			spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
			cal,
		), nil

	case "Ходьба", "ходьба":
		cal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			spentenergy.Distance(t.Steps, t.Height),
			spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
			cal,
		), nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
