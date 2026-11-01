package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	split := strings.Split(datastring, ",")
	if len(split) != 3 {
		return fmt.Errorf("Количество запятых не равно 2\n")
	}
	steps, err := strconv.Atoi(split[0])
	if err != nil {
		return fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if steps < 1 {
		return fmt.Errorf("Кол-во шагов должно быть положительным\n")
	}
	t.Steps = steps
	t.TrainingType = split[1]
	duration, err := time.ParseDuration(split[2])
	if err != nil {
		return fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Длительность должна быть больше нуля\n")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		cals, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "Ошибка при расчете калорий при беге", err
		}
		durInH := t.Duration.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durInH, distance, meanSpeed, cals), nil
	case "Ходьба":
		cals, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "Ошибка при расчете калорий при беге", err
		}
		durInH := t.Duration.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durInH, distance, meanSpeed, cals), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
