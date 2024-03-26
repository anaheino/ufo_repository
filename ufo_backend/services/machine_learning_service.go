package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"ufo_backend/structs"
)

type MachineLearningService struct {
	randomForestURL string
}

func NewMachineLearningService(randomForestURL string) *MachineLearningService {
	return &MachineLearningService{randomForestURL: randomForestURL}
}

func (s *MachineLearningService) Initialize() error {
	pythonBackEndURL := fmt.Sprintf("%s", os.Getenv("MACHINE_LEARNING_URI"))
	if pythonBackEndURL == "" {
		pythonBackEndURL = "http://localhost:5000"
	}
	s.randomForestURL = pythonBackEndURL + "/coordinates/likelihood"
	return nil
}

func (s *MachineLearningService) RandomForestCoordinates(jsonData []byte) structs.Probability {
	client := &http.Client{}
	req, err := http.NewRequest("POST", s.randomForestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Failure on post request")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	jsonData, err = io.ReadAll(resp.Body)
	var requestBody structs.Probability
	if err := json.Unmarshal(jsonData, &requestBody); err != nil {
	}
	return requestBody
}
