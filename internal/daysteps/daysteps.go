package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("Неверный формат данны: ожидается 2 значения")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать кол-во шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("кол-во шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать длительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть больше 0")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
