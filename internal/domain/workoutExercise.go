package domain

import (
	"errors"

	"github.com/google/uuid"
)

type WorkoutExercise struct {
	ID         uuid.UUID `json:"id"`
	WorkoutID  uuid.UUID `json:"workout_id"`
	ExerciseID uuid.UUID `json:"exercise_id"`
	Sets       int       `json:"sets"`
	Reps       int       `json:"reps"`
	Weight     float64   `json:"weight"`
}

func NewWorkoutExercise(workoutID, exerciseID uuid.UUID, sets, reps int, weight float64) (*WorkoutExercise, error) {
	if sets < 1 {
		return nil, errors.New("sets must be at least 1")
	}
	if reps < 1 {
		return nil, errors.New("reps must be at least 1")
	}
	if weight <= 0 {
		return nil, errors.New("weight must be greater than 0")
	}

	return &WorkoutExercise{
		ID:         uuid.New(),
		WorkoutID:  workoutID,
		ExerciseID: exerciseID,
		Sets:       sets,
		Reps:       reps,
		Weight:     weight,
	}, nil
}
