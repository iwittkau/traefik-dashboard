package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	dashboard "github.com/iwittkau/traefik-dashboard"
	"github.com/containous/traefik/types"
)

var (
	_ dashboard.TraefikClient = &TraefikClient{}
)

type TraefikClient struct {
	host string
}

func NewTraefikClient(host string) *TraefikClient {
	return &TraefikClient{
		host: host,
	}
}

func (c *TraefikClient) GetFrontends() (map[string]types.Frontend, error) {
	u := url.URL{
		Scheme: "http",
		Host:   c.host,
		Path:   "api/providers/docker/frontends",
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]types.Frontend

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
