package logging

import (
	"github.com/sirupsen/logrus"
)

type Error struct {
	log        *logrus.Logger
	err        error
	key        string
	httpStatus int
}

func NewError(log *logrus.Logger, err error, key string, httpStatus int) Error {
	return Error{
		log:        log,
		err:        err,
		key:        key,
		httpStatus: httpStatus,
	}
}

func (e Error) Log(msg string) {
	e.log.WithFields(logrus.Fields{
		"key":         e.key,
		"error":       e.err.Error(),
		"http_status": e.httpStatus,
	}).Error(msg)
}
