package models

import (
	"time"

	"github.com/google/uuid"
)

// WorkoutType rappresenta la tipologia di allenamento definita dal personal trainer.
type WorkoutType struct {
	CustomModel
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// TrainingPlan rappresenta il piano di allenamento, associato ad un tipo specifico.
type TrainingPlan struct {
	CustomModel
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	StartDate   time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate     time.Time `gorm:"column:end_date" json:"endDate"`

	UserID uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`

	// Associa il piano ad una tipologia di allenamento
	WorkoutTypeID uuid.UUID   `gorm:"type:uuid;not null" json:"workoutTypeId"`
	WorkoutType   WorkoutType `gorm:"foreignKey:WorkoutTypeID" json:"workoutType"`

	Workouts []Workout `gorm:"foreignKey:TrainingPlanID" json:"workouts"`
}

// Workout rappresenta la scheda di allenamento creata per l'utente.
type Workout struct {
	CustomModel
	Name string `gorm:"column:name;not null" json:"name"`

	TrainingPlanID uuid.UUID    `gorm:"type:uuid;not null" json:"trainingPlanId"`
	TrainingPlan   TrainingPlan `gorm:"foreignKey:TrainingPlanID" json:"trainingPlan"`

	// Utilizziamo una tabella di join esplicita per associare gli esercizi con i dettagli della scheda.
	WorkoutExercises []WorkoutExercise `gorm:"foreignKey:WorkoutID" json:"workoutExercises"`
}

// MuscleGroup organizza gli esercizi per gruppo muscolare.
type MuscleGroup struct {
	CustomModel
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description" json:"description"`

	Exercises []Exercise `gorm:"foreignKey:MuscleGroupID" json:"exercises"`
}

// Exercise rappresenta l’esercizio base, che poi verrà arricchito nei workout.
type Exercise struct {
	CustomModel
	Name string `gorm:"column:name;not null" json:"name"`

	MuscleGroupID uuid.UUID   `gorm:"type:uuid;not null" json:"muscleGroupId"`
	MuscleGroup   MuscleGroup `gorm:"foreignKey:MuscleGroupID" json:"muscleGroup"`
}

// WorkoutExercise rappresenta la scheda dell'esercizio definita dal personal trainer,
// includendo dettagli come numero di serie, ripetizioni, durata e note.
type WorkoutExercise struct {
	CustomModel
	WorkoutID uuid.UUID `gorm:"type:uuid;not null" json:"workoutId"`
	Workout   Workout   `gorm:"foreignKey:WorkoutID" json:"workout"`

	ExerciseID uuid.UUID `gorm:"type:uuid;not null" json:"exerciseId"`
	Exercise   Exercise  `gorm:"foreignKey:ExerciseID" json:"exercise"`

	// Campi della scheda (template) dell'esercizio
	Sets        int    `json:"sets"`
	Repetitions int    `json:"repetitions"`
	Duration    int    `json:"duration"` // ad es. in secondi o minuti
	Notes       string `json:"notes"`
}

// WorkoutProgress registra la sessione di allenamento eseguita dall'utente.
type WorkoutProgress struct {
	CustomModel
	WorkoutID uuid.UUID `gorm:"type:uuid;not null" json:"workoutId"`
	Workout   Workout   `gorm:"foreignKey:WorkoutID" json:"workout"`

	UserID uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`

	// Data della sessione di allenamento
	Date time.Time `gorm:"column:date" json:"date"`

	// Progresso per ciascun esercizio del workout
	ExerciseProgress []ExerciseProgress `gorm:"foreignKey:WorkoutProgressID" json:"exerciseProgress"`
}

// ExerciseProgress registra i progressi reali dell'utente per un determinato esercizio,
// facendo riferimento alla scheda (WorkoutExercise) definita dal personal trainer.
type ExerciseProgress struct {
	CustomModel
	WorkoutProgressID uuid.UUID       `gorm:"type:uuid;not null" json:"workoutProgressId"`
	WorkoutProgress   WorkoutProgress `gorm:"foreignKey:WorkoutProgressID" json:"workoutProgress"`

	// Riferimento alla scheda dell'esercizio definita nel workout
	WorkoutExerciseID uuid.UUID       `gorm:"type:uuid;not null" json:"workoutExerciseId"`
	WorkoutExercise   WorkoutExercise `gorm:"foreignKey:WorkoutExerciseID" json:"workoutExercise"`

	// Dettagli relativi ai progressi registrati dall'utente
	ActualRepetitions int     `json:"actualRepetitions"`
	ActualWeight      float64 `json:"actualWeight"`
	ActualDuration    int     `json:"actualDuration"`
	Notes             string  `json:"notes"`
}
