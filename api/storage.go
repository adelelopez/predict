package api

type Storage interface {
	SavePrediction(p *Prediction) error
	UpdatePrediction(p *Prediction) (*Prediction, error)
	GetPrediction(id string) (*Prediction, error)
	GetPredictions(p *Prediction) ([]Prediction, error)
}
