package catalog

import (
	"encoding/json"
	"fmt"
	"net/http"

	profilesv1 "github.com/weaveworks/profiles/api/v1alpha1"
)

// Search queries the catalog at catalogURL for profiles matching the provided searchName.
// If no searchname is provided it returns all profiles
func Search(catalogClient CatalogClient, searchName string) ([]profilesv1.ProfileCatalogEntry, error) {
	var q map[string]string

	if searchName != "" {
		q = map[string]string{
			"name": searchName,
		}
	}
	data, statusCode, err := catalogClient.DoRequest("/profiles", q)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch catalog: %w", err)
	}

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch profile from catalog, status code %d", statusCode)
	}
	var profiles []profilesv1.ProfileCatalogEntry
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, fmt.Errorf("failed to parse catalog: %w", err)
	}

	return profiles, nil
}
