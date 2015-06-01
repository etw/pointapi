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

// PointAPI object
type PointAPI struct {
	httpClient *http.Client
	auth       *string
}

// Returns new PointAPI object
func New(c *http.Client, a *string) *PointAPI {
	api := new(PointAPI)
	api.httpClient = c
	api.auth = a
	return api
}

// Private metamethod for talking to point.im API
func (api *PointAPI) metaGet(s *string) (*PostList, error) {
	var resmap PostList
	resp, err := api.httpClient.Get(fmt.Sprint(apiPrefix, *s))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
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
