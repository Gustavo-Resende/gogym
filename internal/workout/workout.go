package workout

import (
	"errors"

	"github.com/Gustavo-Resende/gogym/internal/exercise"
)

type Workout struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Exercises   []exercise.Exercise `json:"exercises"`
}

func NewWorkout(name string, description string) (*Workout, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &Workout{
		Name:        name,
		Description: description,
	}, nil
}

func (w *Workout) AddExercise(exercise exercise.Exercise) error {
	for _, e := range w.Exercises {
		if e.ID == exercise.ID {
			return errors.New("exercise already exists")
		}
	}
	w.Exercises = append(w.Exercises, exercise)
	return nil
}
