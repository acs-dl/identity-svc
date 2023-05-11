package handlers

import (
	"net/http"

	"github.com/acs-dl/identity-svc/resources"
	"gitlab.com/distributed_lab/ape"
)

func GetPositions(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, resources.PositionsResponse{
		Data: resources.Positions{
			Key: resources.Key{
				ID:   "1",
				Type: resources.POSITIONS,
			},
			Attributes: resources.PositionsAttributes{
				Positions: Positions(r),
			},
		}})
}
