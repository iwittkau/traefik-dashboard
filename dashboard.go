package dashboard

import "github.com/containous/traefik/types"

type TraefikClient interface {
	GetFrontends() (map[string]types.Frontend, error)
}
