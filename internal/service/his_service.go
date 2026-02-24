package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HISService struct {
	BaseURL string
	Client  *http.Client
}

func NewHISService() *HISService {
	return &HISService{
		BaseURL: "https://hospital-a.api.co.th",
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type HISResponse struct {
	FirstNameEN string `json:"first_name_en"`
	LastNameEN  string `json:"last_name_en"`
	NationalID  string `json:"national_id"`
	PassportID  string `json:"passport_id"`
	Gender      string `json:"gender"`
}

func (h *HISService) FetchPatient(id string) (*HISResponse, error) {
	url := fmt.Sprintf("%s/patient/search/%s", h.BaseURL, id)

	resp, err := h.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("patient not found in HIS")
	}

	var result HISResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}