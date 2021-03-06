package goieeeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrNilClient = errors.New("given client is nil")
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

type UserInfo struct {
	Institute                     bool    `json:"institute"`
	Member                        bool    `json:"member"`
	Individual                    bool    `json:"individual"`
	Guest                         bool    `json:"guest"`
	SubscribedContent             bool    `json:"subscribedContent"`
	FileCabinetContent            bool    `json:"fileCabinetContent"`
	FileCabinetUser               bool    `json:"fileCabinetUser"`
	InstitutionalFileCabinetUser  bool    `json:"institutionalFileCabinetUser"`
	IP                            *string `json:"ip,omitempty"`
	ShowPatentCitations           bool    `json:"showPatentCitations"`
	ShowGet802Link                bool    `json:"showGet802Link"`
	ShowOpenURLLink               *bool   `json:"showOpenUrlLink,omitempty"`
	Tracked                       *bool   `json:"tracked,omitempty"`
	DelegatedAdmin                *bool   `json:"delegatedAdmin,omitempty"`
	Desktop                       *bool   `json:"desktop,omitempty"`
	IsInstitutionDashboardEnabled *bool   `json:"isInstitutionDashboardEnabled,omitempty"`
	IsInstitutionProfileEnabled   *bool   `json:"isInstitutionProfileEnabled,omitempty"`
	IsRoamingEnabled              *bool   `json:"isRoamingEnabled,omitempty"`
	IsDelegatedAdmin              *bool   `json:"isDelegatedAdmin,omitempty"`
	IsMdl                         *bool   `json:"isMdl,omitempty"`
	IsCwg                         *bool   `json:"isCwg,omitempty"`
}

type ErrUnexpectedStatus struct {
	StatusCode int
	Body       []byte
}

func (e ErrUnexpectedStatus) Error() string {
	return fmt.Sprintf("got unexpected status code %d with body %s", e.StatusCode, e.Body)
}

var _ error = (*ErrUnexpectedStatus)(nil)

func getEndp(client HTTPClient, id int, endp string, dst interface{}) error {
	if client == nil {
		return ErrNilClient
	}

	// Issue the API call
	url := fmt.Sprintf("https://ieeexplore.ieee.org/rest/document/%d/%s", id, endp)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Referer", "https://ieeexplore.ieee.org")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Extract the response body
	body, _ := io.ReadAll(res.Body)

	// Check status code
	if res.StatusCode != http.StatusOK {
		return &ErrUnexpectedStatus{
			StatusCode: res.StatusCode,
			Body:       body,
		}
	}

	// Unmarshal JSON REST response
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
