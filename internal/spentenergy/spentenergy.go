package spentenergy

import (
	"errors"
	"log"
	"time"
)

var errInvalidInput = errors.New("неверные входные данные")

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		log.Println("количество шагов должно быть больше 0")
		return 0, errInvalidInput
	}
	if weight <= 0 {
		log.Println("вес должен быть больше 0")
		return 0, errInvalidInput
	}
	if height <= 0 {
		log.Println("рост должен быть больше 0")
		return 0, errInvalidInput
	}
	if duration <= 0 {
		log.Println("продолжительность должна быть больше 0")
		return 0, errInvalidInput
	}
	speed := MeanSpeed(steps, height, duration)
	return walkingCaloriesCoefficient * weight * speed * duration.Hours(), nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		log.Println("количество шагов должно быть больше 0")
		return 0, errInvalidInput
	}
	if weight <= 0 {
		log.Println("вес должен быть больше 0")
		return 0, errInvalidInput
	}
	if height <= 0 {
		log.Println("вес должен быть больше 0")
		return 0, errInvalidInput
	}
	if duration <= 0 {
		log.Println("продолжительность должна быть больше 0")
		return 0, errInvalidInput
	}

	speed := MeanSpeed(steps, height, duration)
	return weight * speed * duration.Minutes() / minInH, nil
}
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	speed := dist / duration.Hours()
	return speed
}

func Distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	return float64(steps) * lenStep / mInKm
}
