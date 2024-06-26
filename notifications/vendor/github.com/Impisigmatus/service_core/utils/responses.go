package utils

import (
	"fmt"
	"net/http"

	"github.com/Impisigmatus/service_core/log"
	jsoniter "github.com/json-iterator/go"
)

func WriteNoContent(w http.ResponseWriter) {
	WriteString(w, http.StatusNoContent, nil, "")
}

func WriteString(w http.ResponseWriter, status int, err error, str string, args ...any) {
	const (
		header  = "Content-Type"
		content = "text/plain"
	)

	if err != nil {
		log.Error("", err)
	}

	w.Header().Set(header, content)
	w.WriteHeader(status)

	if _, err := w.Write([]byte(fmt.Sprintf(str, args...))); err != nil {
		log.Error("Invalid write response body", err)
	}
}

func WriteObject(w http.ResponseWriter, obj interface{}) {
	const (
		header  = "Content-Type"
		content = "application/json"
	)

	w.Header().Set(header, content)
	w.WriteHeader(http.StatusOK)

	data, err := jsoniter.Marshal(obj)
	if err != nil {
		WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if _, err := w.Write(data); err != nil {
		log.Error("Invalid write response body", err)
	}
}
