package api

import "time"

type Prediction struct {
	ID          string
	Name        string
	Probability float64
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

type History struct {
	Predictions []Prediction
}

func CreatePrediction(p Prediction, s Storage) (*Prediction, error) {
	err := s.SavePrediction(&p)
	return &p, err
}

func JudgePrediction(p Prediction, s Storage) (*Prediction, error) {
	return nil, nil
}

func GetStats(s Storage) (*Statistics, error) {
	return nil, nil
}

func GetHistory(s Storage) (*History, error) {
	return nil, nil
}
