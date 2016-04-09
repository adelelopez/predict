package api

type Storage interface {
	SavePrediction(p *Prediction) (*Prediction, error)
	UpdatePrediction(id string, p Prediction) error
	GetPrediction(id string) (*Prediction, error)
	GetPredictions(p *Prediction) ([]Prediction, error)
}
