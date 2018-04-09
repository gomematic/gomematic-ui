package handler

import (
	"fmt"
	"net/http"

	"github.com/gomematic/gomematic-ui/pkg/assets"
	"github.com/gomematic/gomematic-ui/pkg/config"
)

// Static handles the delivery of all static assets.
func Static(cfg *config.Config) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf(
			"%sassets",
			cfg.Server.Root,
		),
		http.FileServer(
			assets.Load(cfg),
		),
	)
}
