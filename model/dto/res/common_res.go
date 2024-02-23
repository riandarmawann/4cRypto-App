package res

import (
	"4crypto/utils/model_util"
)

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}

// paged response
type PagedResponse struct {
	Status Status            `json:"status"`
	Data   any               `json:"data"`
	Paging model_util.Paging `json:"paging"`
}

type CommonResponse struct {
	Code    int
	Status  string
	Message string
	Data    any
}