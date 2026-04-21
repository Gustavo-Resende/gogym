package exercise

import "errors"

type Exercise struct {
	ID          int         `json:"id"`
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
		Name:        name,
		Description: description,
		MuscleGroup: muscleGroup,
	}, nil
}
