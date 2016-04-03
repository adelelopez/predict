package filestorage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/adelethalialopez/predict/api"
)

type FileStorage struct {
	Filename string
}

func (fs FileStorage) SavePrediction(p *api.Prediction) error {
	file, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		file, err = os.Create(fs.Filename)
		fmt.Printf("History file not found, so creating: %s\n", fs.Filename)
	}

	defer file.Close()

	if p.CreatedAt == nil {
		now := time.Now()
		p.CreatedAt = &now
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (fs FileStorage) UpdatePrediction(p *api.Prediction) (*api.Prediction, error) {
	return nil, nil
}

func (fs FileStorage) GetPrediction(id string) (*api.Prediction, error) {
	return nil, nil
}

func (fs FileStorage) GetPredictions(p *api.Prediction) ([]api.Prediction, error) {
	return nil, nil
}
