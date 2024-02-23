package res

import "4crypto/model/entity"

var ResponseExhange struct {
	Status entity.CmcExchangeStatus `json:"status"`
	Data   []entity.CmcExhange      `json:"data"`
}
