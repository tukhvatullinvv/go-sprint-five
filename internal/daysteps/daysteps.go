package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	split := strings.Split(datastring, ",")
	if len(split) != 2 {
		return fmt.Errorf("В строке отсутствуют запятые или их больше 1\n")
	}
	steps, err := strconv.Atoi(split[0])
	if err != nil {
		return fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if steps < 1 {
		return fmt.Errorf("Кол-во шагов должно быть положительным\n")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(split[1])
	if err != nil {
		return fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть больше нуля\n")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	ccals, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, ccals), nil
}
