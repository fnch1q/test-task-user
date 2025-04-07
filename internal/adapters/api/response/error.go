package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"test-task-user/internal/adapters/api/logging"
	errorStatus "test-task-user/internal/errors"

	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidInput = errors.New("invalid_input")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     messages,
	}
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}

func NewErrorWithErrorStatus(err error, w http.ResponseWriter, log *logrus.Logger, logKey, message string) {
	re, ok := err.(*errorStatus.ErrorStatus)
	var statusCode int
	if ok {
		statusCode = re.StatusCode
	} else {
		statusCode = http.StatusInternalServerError
	}
	NewError(err, statusCode).Send(w)

	logMsg := message
	logging.NewError(log, err, logKey, statusCode).Log(logMsg)
}
