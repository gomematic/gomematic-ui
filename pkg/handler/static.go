package handler

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gomematic/gomematic-ui/pkg/assets"
	"github.com/gomematic/gomematic-ui/pkg/config"
)

// Static handles the delivery of all static assets.
func Static(logger log.Logger) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf(
			"%sassets",
			config.Server.Root,
		),
		http.FileServer(
			assets.Load(logger),
		),
	)
}
