package http_test

import (
	"testing"

	"github.com/iwittkau/traefik-dashboard/http"
)

func Test_TraefikClient(t *testing.T) {
	c := http.NewTraefikClient("traefik.docker.localhost")

	res, err := c.GetFrontends()

	if err != nil {
		t.Error(err.Error())
	}

	t.Log(res)
}
