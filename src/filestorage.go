package filestorage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/adelethalialopez/predict/api"
	"github.com/pborman/uuid"
)

type FileStorage struct {
	Filename string
	data     []api.Prediction
}

func (fs *FileStorage) readData() error {
	fs.data = make([]api.Prediction, 0, len(fs.data))

	file, err := os.OpenFile(fs.Filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("History file not found, so creating: %s\n", fs.Filename)
		file, err = os.Create(fs.Filename)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	for {
		var np api.Prediction
		if err := dec.Decode(&np); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fs.data = append(fs.data, np)
	}
	return nil
}

func (fs *FileStorage) writeData() error {
	dataBytes := make([][]byte, 0, len(fs.data))

	for _, p := range fs.data {
		bytes, err := json.Marshal(p)
		if err != nil {
			return err
		}

		dataBytes = append(dataBytes, bytes)
	}

	err := ioutil.WriteFile(fs.Filename, bytes.Join(dataBytes, []byte("\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}

func filter(filter *api.Prediction, p api.Prediction) bool {
	if filter == nil {
		return true
	}
	if filter.ID != "" && filter.ID != p.ID {
		return false
	}
	if filter.Name != "" && filter.Name != p.Name {
		return false
	}
	if filter.Outcome != nil && filter.Outcome != p.Outcome {
		return false
	}
	if filter.Probability != nil && filter.Probability != p.Probability {
		return false
	}
	if filter.CreatedAt != nil && filter.CreatedAt != p.CreatedAt {
		return false
	}
	// TODO: filter tags
	return true
}

func (fs *FileStorage) SavePrediction(p *api.Prediction) (*api.Prediction, error) {
	file, err := os.OpenFile(fs.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		file, err = os.Create(fs.Filename)
		fmt.Printf("History file not found, so creating: %s\n", fs.Filename)
	}

	defer file.Close()

	p.ID = uuid.NewRandom().String()

	if p.CreatedAt == nil {
		now := time.Now()
		p.CreatedAt = &now
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	_, err = file.Write(bytes)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (fs *FileStorage) UpdatePrediction(id string, p api.Prediction) (*api.Prediction, error) {
	err := fs.readData()
	if err != nil {
		return nil, err
	}

	ps := make([]api.Prediction, 0, len(fs.data))

	returnP := p
	for _, np := range fs.data {
		if np.ID == id {
			p.ID = id

			if p.CreatedAt == nil {
				p.CreatedAt = np.CreatedAt
			}
			if p.Name == "" {
				p.Name = np.Name
			}
			if p.Probability == nil {
				p.Probability = np.Probability
			}
			if p.Outcome == nil {
				p.Outcome = np.Outcome
			}

			// TODO: don't update tags if blank

			ps = append(ps, p)
			returnP = p
		} else {
			ps = append(ps, np)
		}
	}

	fs.data = ps
	err = fs.writeData()
	if err != nil {
		return nil, err
	}

	return &returnP, nil
}

func (fs *FileStorage) GetPrediction(id string) (*api.Prediction, error) {
	err := fs.readData()
	if err != nil {
		return nil, err
	}

	for _, prediction := range fs.data {
		if prediction.ID == id {
			return &prediction, nil
		}
	}

	return nil, nil
}

func (fs *FileStorage) GetPredictions(p *api.Prediction) ([]api.Prediction, error) {
	err := fs.readData()
	if err != nil {
		return nil, err
	}

	retPredictions := make([]api.Prediction, 0, len(fs.data))
	for _, prediction := range fs.data {
		if filter(p, prediction) {
			retPredictions = append(retPredictions, prediction)
		}
	}
	return retPredictions, nil
}
