package pointapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Prefix for point.im API endpoints
const apiPrefix = "https://point.im/api"

// API object
type API struct {
	httpClient *http.Client
	auth       *string
}

// Returns new PointAPI object
func New(c *http.Client, a *string) *API {
	api := new(API)
	api.httpClient = c
	api.auth = a
	return api
}

// Private metamethod for talking to point.im API
func (api *API) metaGet(s *string) (*PostList, error) {
	var resmap PostList

	req, err := http.NewRequest("GET", fmt.Sprint(apiPrefix, *s), nil)
	if err != nil {
		return nil, err
	}
	if api.auth != nil {
		req.Header.Set("Authorization", *api.auth)
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resmap)
	if err != nil {
		return nil, err
	}
	return &resmap, nil
}
