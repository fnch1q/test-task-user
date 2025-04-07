package logging

import "github.com/sirupsen/logrus"

type Info struct {
	log        *logrus.Logger
	key        string
	httpStatus int
}

func NewInfo(log *logrus.Logger, key string, httpStatus int) Info {
	return Info{
		log:        log,
		key:        key,
		httpStatus: httpStatus,
	}
}

func (i Info) Log(msg string) {
	i.log.WithFields(logrus.Fields{
		"key":         i.key,
		"http_status": i.httpStatus,
	}).Infof(msg)
}
