package http

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func newErrorResponse(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		logrus.Error(err)
	}
	customErr := CustomError{
		Message: message,
		Code:    code,
	}
	w.WriteHeader(customErr.Code)
	json.NewEncoder(w).Encode(customErr)
}
