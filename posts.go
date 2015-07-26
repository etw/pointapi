package pointapi

import (
	"fmt"
	"text/template"
	"strings"
)

// Method to get posts form https://point.im/all
func (api *API) GetAll(before int) (*PostList, error) {
	path := fmt.Sprintf("/all?before=%d", before)
	return api.metaGet(&path)
}

// Method to get posts filtered by tags
func (api *API) GetTags(before int, tags []string) (*PostList, error) {
	var esctags []string
	for _, v := range tags {
		esctags = append(esctags, template.URLQueryEscaper(v))
	}
	path := fmt.Sprintf("/tags?before=%d&tag=%s", before, strings.Join(esctags, "&tag="))
	return api.metaGet(&path)
}
