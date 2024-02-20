package res

import "4crypto/model/entity"

var ResponseFiat struct {
	Status entity.CmcFiatStatus `json:"status"`
	Data   []entity.CmcFiat     `json:"data"`
}
