package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error){
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	
	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResp, nil
}
