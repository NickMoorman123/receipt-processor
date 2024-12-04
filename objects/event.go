package objects

import (
	"encoding/json"
	"net/http"
)

type GetRequest struct {
	UUID string `json:"id"`
}

type GetResponseWrapper struct {
	Points	int	`json:"points"`
	Code	int	`json:"-"`
}

func (e *GetResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

func (e *GetResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}
	return e.Code
}

type ProcessRequest struct {
	Receipt *Receipt `json:"receipt"`
}

type ProcessResponseWrapper struct {
	UUID	string	`json:"id"`
	Code	int		`json:"-"`
}

func (e *ProcessResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

func (e *ProcessResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}
	return e.Code
}