package action

import (
	"net/http"
	"strconv"

	"test-task-user/internal/adapters/api/logging"
	"test-task-user/internal/adapters/api/response"
	"test-task-user/internal/usecase"

	"github.com/sirupsen/logrus"
)

type FindUserByIDAction struct {
	uc  usecase.FindUserByIDUseCase
	log *logrus.Logger
}

func NewFindUserByIDAction(uc usecase.FindUserByIDUseCase, log *logrus.Logger) FindUserByIDAction {
	return FindUserByIDAction{
		uc:  uc,
		log: log,
	}
}

func (a FindUserByIDAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_user_by_id"

	var input usecase.FindUserByIDInput
	var keysStr = r.URL.Query()
	id, err := strconv.Atoi(keysStr["id"][0])
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("invalid_id_format")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	input.ID = uint32(id)

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewErrorWithErrorStatus(err, w, a.log, logKey, "error when find user by id")
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success find user by id")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
