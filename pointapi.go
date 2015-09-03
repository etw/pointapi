package pointapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Prefix for point.im API endpoints
const (
	POINTIM  = "https://point.im"
	POINTAPI = "https://point.im/api"
)

// API object
type API struct {
	httpClient *http.Client
	prefix     string
	auth       *string
}

// Returns new PointAPI object
func New(c *http.Client, p string, a *string) *API {
	api := new(API)
	api.httpClient = c
	api.prefix = p
	api.auth = a
	return api
}

// Private metamethod for talking to point.im API
func (api *API) metaGet(s *string) (p *PostList, e error) {
	p = new(PostList)

	if req, err := http.NewRequest("GET", fmt.Sprint(api.prefix, *s), nil); err != nil {
		return nil, err
	} else {
		if api.auth != nil {
			req.Header.Set("Authorization", *api.auth)
		}

		if resp, err := api.httpClient.Do(req); err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return nil, errors.New(resp.Status)
			}

			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return nil, err
			} else {
				if err := json.Unmarshal(body, p); err != nil {
					return nil, err
				}
			}
		}
	}

	return p, nil
}
