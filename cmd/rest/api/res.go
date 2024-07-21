package api

import (
	j "encoding/json"
	"errors"
	"net/http"

	"awesomeProject/domain"
	"awesomeProject/lib/logger"
)

func jsonRes(w http.ResponseWriter, statusCode int, out any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	marshal, err := j.Marshal(out)
	if err != nil {
		return
	}

	_, _ = w.Write(marshal)
}

func errorRes(err error, res *http.ResponseWriter) {
	// if client safe error, return it

	logger.New(false).Error(err)

	e := fromError(err)

	jsonRes(*res, e.Status, map[string]string{"error": e.Message})
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func fromError(err error) Error {
	var (
		apiErr Error
		svcErr *domain.Error
	)

	if errors.As(err, &svcErr) {
		apiErr.Message = svcErr.AppErr().Error()
		switch {
		case errors.Is(svcErr, domain.ErrBadRequest):
			apiErr.Status = http.StatusBadRequest
		case errors.Is(svcErr, domain.ErrNotFound):
			apiErr.Status = http.StatusNotFound
		case errors.Is(svcErr, domain.ErrInternal):
			apiErr.Status = http.StatusInternalServerError
		default:
			apiErr.Status = http.StatusInternalServerError
		}
	}

	return apiErr
}
