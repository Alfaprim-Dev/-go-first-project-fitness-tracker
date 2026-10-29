package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	errInvalidSteps    = errors.New("Количество шагов должно быть больше 0")
	errInvalidWeight   = errors.New("Вес должен быть больше 0")
	errInvalidHeight   = errors.New("Рост должен быть больше 0")
	errInvalidDuration = errors.New("Длительность должна быть больше 0")
)

func validate(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errInvalidSteps
	}
	if weight <= 0 {
		return errInvalidWeight
	}
	if height <= 0 {
		return errInvalidHeight
	}
	if duration <= 0 {
		return errInvalidDuration
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validate(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validate(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	return (weight * meanSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
