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
	return &p, s.UpdatePrediction(p.ID, p)
}

func JudgeLastPrediction(outcome bool, s Storage) (*Prediction, error) {
	ps, err := s.GetPredictions(nil)
	if err != nil {

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
	return &p, s.UpdatePrediction(id, p)
}

func GetStats(s Storage) (*Statistics, error) {
	return nil, nil
}

func GetHistory(s Storage) ([]Prediction, error) {
	return s.GetPredictions(nil)
}
