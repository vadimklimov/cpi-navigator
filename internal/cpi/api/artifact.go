package api

import (
	"fmt"

	"github.com/vadimklimov/cpi-navigator/internal/cpi/client"
)

type IntegrationArtifact struct {
	ID          string `json:"Id"`
	Version     string `json:"Version"`
	PackageID   string `json:"PackageId"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	CreatedBy   string `json:"CreatedBy"`
	CreatedAt   int64  `json:"CreatedAt,string"`
	ModifiedBy  string `json:"ModifiedBy"`
	ModifiedAt  int64  `json:"ModifiedAt,string"`
}

func IntegrationArtifactsByPackageAndType(packageID, artifactType string) ([]IntegrationArtifact, error) {
	var responseBody struct {
		Root struct {
			Results []IntegrationArtifact `json:"results"`
		} `json:"d"`
	}

	var entitySetName string

	supportedArtifactTypes := SupportedArtifactTypes()

	switch artifactType {
	case supportedArtifactTypes.Designtime.IntegrationFlow.Name:
		entitySetName = supportedArtifactTypes.Designtime.IntegrationFlow.EntitySetName
	case supportedArtifactTypes.Designtime.ValueMapping.Name:
		entitySetName = supportedArtifactTypes.Designtime.ValueMapping.EntitySetName
	case supportedArtifactTypes.Designtime.MessageMapping.Name:
		entitySetName = supportedArtifactTypes.Designtime.MessageMapping.EntitySetName
	case supportedArtifactTypes.Designtime.ScriptCollection.Name:
		entitySetName = supportedArtifactTypes.Designtime.ScriptCollection.EntitySetName
	}

	res, err := client.NewClient().R().
		SetResult(&responseBody).
		SetPathParams(map[string]string{
			"package":   packageID,
			"entitySet": entitySetName,
		}).
		SetQueryParam("$format", "json").
		Get("IntegrationPackages('{package}')/{entitySet}")
	if err != nil {
		return nil, fmt.Errorf("error when calling %s: %w", res.Request.URL, err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("error when calling %s: %s", res.Request.URL, res.Status())
	}

	return responseBody.Root.Results, nil
}
