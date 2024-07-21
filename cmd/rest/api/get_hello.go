package api

import (
	"net/http"

	"awesomeProject/internal/usecase"
)

type exampleRes struct {
	Message string `json:"message"`
}

func (s Server) GetHello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorRes(err, &w)
		return
	}

	res, err := usecase.NewSayHello(s.container.ExampleService).Execute(r.Form.Get("name"))
	if err != nil {
		errorRes(err, &w)
		return
	}

	jsonRes(w, http.StatusOK, exampleRes{Message: res})
}
