package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func TrackAddress(w http.ResponseWriter, r *http.Request) {
	ape.RenderErr(w, problems.InternalError())
	//TODO: realize router handler
}
