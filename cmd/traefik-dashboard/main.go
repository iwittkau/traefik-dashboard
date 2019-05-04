//go:generate statik -src=../../public

package main

import (
	"flag"
	"log"

	_ "github.com/iwittkau/traefik-dashboard/cmd/traefik-dashboard/statik"
	"github.com/iwittkau/traefik-dashboard/http"
	"github.com/rakyll/statik/fs"
)

var (
	host                 string
	templatesPath        string
	address              string
	defaultTemplatesPath = "/templates"
	defaultAddress       = ":80"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&host, "h", "", "set the traefik host")
	flag.StringVar(&templatesPath, "t", defaultTemplatesPath, "set the path for html templates")
	flag.StringVar(&address, "a", defaultAddress, "set the listening address")
	flag.Parse()

	c := http.NewTraefikClient(host)

	f := http.NewFrontend(c, templatesPath, statikFS)

	f.Open(address)
}
