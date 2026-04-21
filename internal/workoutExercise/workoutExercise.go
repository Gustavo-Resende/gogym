package workoutExercise

type WorkoutExercise struct {
	GUID         string  `json:"id"`
	WorkoutGUID  string  `json:"workout_id"`
	ExerciseGUID string  `json:"exercise_id"`
	Sets         int     `json:"sets"`
	Reps         int     `json:"reps"`
	Weight       float64 `json:"weight"`
}
