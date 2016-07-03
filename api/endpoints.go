package api

import (
	"time"
)

type Prediction struct {
	ID          string
	Name        string
	Probability *float64
	Outcome     *bool
	Tags        []string
	CreatedAt   *time.Time
}

type Bucket struct {
	Score      float64
	LeftBound  float64
	RightBound float64
	Star       float64
	Mean       float64
	Size       int
}

type Statistics struct {
	Score            float64
	TotalPredictions int
	NumberOfBuckets  int
	Buckets          []Bucket
}

var (
	True  = true
	False = false
)

func CreatePrediction(p Prediction, s Storage) (*Prediction, error) {
	return s.SavePrediction(&p)
}

func UpdatePrediction(p Prediction, s Storage) (*Prediction, error) {
	return s.UpdatePrediction(p.ID, p)
}

func AmendLastPrediction(amendments Prediction, s Storage) (*Prediction, error) {
	ps, err := s.GetPredictions(nil)
	if err != nil {
		return nil, err
	}

	mostRecent := time.Unix(0, 0)

	id := ""
	for _, prediction := range ps {
		if (*prediction.CreatedAt).After(mostRecent) {
			mostRecent = *prediction.CreatedAt
			id = prediction.ID
		}
	}

	return s.UpdatePrediction(id, amendments)
}

func JudgeLastPrediction(outcome bool, s Storage) (*Prediction, error) {
	ps, err := s.GetPredictions(nil)
	if err != nil {
		return nil, err
	}

	mostRecent := time.Unix(0, 0)

	id := ""
	for _, prediction := range ps {
		if (*prediction.CreatedAt).After(mostRecent) && prediction.Outcome == nil {
			mostRecent = *prediction.CreatedAt
			id = prediction.ID
		}
	}

	p := Prediction{
		Outcome: &outcome,
	}
	return s.UpdatePrediction(id, p)
}

func GetStats(s Storage) (*Statistics, error) {
	hist, err := s.GetPredictions(nil)
	if err != nil {
		return nil, err
	}

	// Decide on buckets
	// Add predictions to their bucket
	// Calculate total score, and scores for each bucket

	stats := Statistics{
		TotalPredictions: len(hist),
	}
	return &stats, nil
}

func GetHistory(s Storage) ([]Prediction, error) {
	return s.GetPredictions(nil)
}
