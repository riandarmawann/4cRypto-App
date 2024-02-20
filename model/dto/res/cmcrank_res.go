package res

import "4crypto/model/entity"

var Response struct {
	Status entity.CmcRankStatus `json:"status"`
	Data   []entity.CmcRank     `json:"data"`
}
