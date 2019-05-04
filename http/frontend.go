package http

import (
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	dashboard "github.com/iwittkau/traefik-dashboard"
)

type Frontend struct {
	client        dashboard.TraefikClient
	templatesPath string
	fs            http.FileSystem
}

func NewFrontend(c dashboard.TraefikClient, path string, fs http.FileSystem) *Frontend {
	return &Frontend{
		client:        c,
		templatesPath: path,
		fs:            fs,
	}
}

func (f *Frontend) Open(addr string) error {
	r := gin.Default()

	r.LoadHTMLFiles(f.templatesPath + string(os.PathSeparator) + "index.html")

	r.GET("/", f.handleIndex)
	r.StaticFS("/static", f.fs)

	return r.Run(addr)
}

func (f *Frontend) handleIndex(c *gin.Context) {

	fes, err := f.client.GetFrontends()
	links := []string{}
	for i := range fes {
		for _, v := range fes[i].Routes {
			if strings.HasPrefix(v.Rule, "Host:") {
				links = append(links, strings.TrimPrefix(v.Rule, "Host:"))
			}
		}
	}

	sort.Strings(links)

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"Links": links,
		"Error": err,
	})
}
