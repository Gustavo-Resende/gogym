package domain

import (
	"errors"

	"github.com/google/uuid"
)

type MuscleGroup string

const (
	Chest     MuscleGroup = "chest"
	Triceps   MuscleGroup = "triceps"
	Biceps    MuscleGroup = "biceps"
	Back      MuscleGroup = "back"
	Legs      MuscleGroup = "legs"
	Shoulders MuscleGroup = "shoulders"
	Abs       MuscleGroup = "abs"
)

type Exercise struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	MuscleGroup MuscleGroup `json:"muscle_group"`
}

func NewExercise(name string, description string, muscleGroup MuscleGroup) (*Exercise, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}
	if muscleGroup == "" {
		return nil, errors.New("muscle group is required")
	}

	return &Exercise{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		MuscleGroup: muscleGroup,
	}, nil
}
