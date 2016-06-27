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
	BrierScore float64
	LeftBound  float64
	RightBound float64
	Star       float64
	Mean       float64
	Size       int64
}

type Statistics struct {
	BrierScore       float64
	TotalPredictions int64
	NumberOfBuckets  int64
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
	// TODO: actually write this
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
	return stats, nil
}

func GetHistory(s Storage) ([]Prediction, error) {
	return s.GetPredictions(nil)
}
