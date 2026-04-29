package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Workout struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exercises   []WorkoutExercise `json:"exercises"`
}

func NewWorkout(name string, description string) (*Workout, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &Workout{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}, nil
}

func (w *Workout) AddExercise(we *WorkoutExercise) error {
	for _, e := range w.Exercises {
		if e.ExerciseID == we.ExerciseID {
			return errors.New("exercise already exists")
		}
	}
	w.Exercises = append(w.Exercises, *we)
	return nil
}
