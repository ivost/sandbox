package api

import (
	"net/http"

	"github.com/braincorp/boss_pod_template/pkg/version"
)

// Version godoc
// @Summary Version
// @Description returns podinfo version and git commit hash
// @Tags HTTP API
// @Produce json
// @Router /version [get]
// @Success 200 {object} api.MapResponse
func (s *Server) handleVersion() http.HandlerFunc {
	// this executes once
	result := map[string]string{
		"version": version.VERSION,
		"commit":  version.REVISION,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.JSONResponse(w, r, result)
	}
}

/*
func (s *server) handleSomething() http.HandlerFunc {
    thing := prepareThing()
    return func(w http.ResponseWriter, r *http.Request) {
        // use thing
    }
}
*/
