package api

import (
	"fmt"

	"github.com/vadimklimov/cpi-navigator/internal/cpi/client"
)

type ContentPackage struct {
	ID           string `json:"Id"`
	Version      string `json:"Version"`
	Name         string `json:"Name"`
	ShortText    string `json:"ShortText"`
	Description  string `json:"Description"`
	Vendor       string `json:"Vendor"`
	Mode         string `json:"Mode"`
	CreatedBy    string `json:"CreatedBy"`
	CreationDate int64  `json:"CreationDate,string"`
	ModifiedBy   string `json:"ModifiedBy"`
	ModifiedDate int64  `json:"ModifiedDate,string"`
}

func ContentPackages() ([]ContentPackage, error) {
	var responseBody struct {
		Root struct {
			Results []ContentPackage `json:"results"`
		} `json:"d"`
	}

	res, err := client.NewClient().R().
		SetResult(&responseBody).
		SetQueryParam("$format", "json").
		Get("IntegrationPackages")
	if err != nil {
		return nil, fmt.Errorf("error when calling %s: %w", res.Request.URL, err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("error when calling %s: %s", res.Request.URL, res.Status())
	}

	return responseBody.Root.Results, nil
}
