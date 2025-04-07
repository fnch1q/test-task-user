package action

import (
	"fmt"
	"net/http"
	"strconv"
	"test-task-user/internal/adapters/api/logging"
	"test-task-user/internal/adapters/api/response"
	"test-task-user/internal/usecase"

	"github.com/sirupsen/logrus"
)

type DeleteUserAction struct {
	uc  usecase.DeleteUserUseCase
	log *logrus.Logger
}

func NewDeleteUserAction(uc usecase.DeleteUserUseCase, log *logrus.Logger) DeleteUserAction {
	return DeleteUserAction{
		uc:  uc,
		log: log,
	}
}

func (f DeleteUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "delete_user"

	var keysStr = r.URL.Query()
	id, err := strconv.Atoi(keysStr["id"][0])
	if err != nil {
		logging.NewError(
			f.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("invalid_id_format")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	var input usecase.DeleteUserInput
	input.ID = uint32(id)

	err = f.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewErrorWithErrorStatus(err, w, f.log, logKey, fmt.Sprintf("error when %s", logKey))
		return
	}

	logging.NewInfo(f.log, logKey, http.StatusOK).
		Log(fmt.Sprintf("success when %s", logKey))

	response.NewSuccess(nil, http.StatusOK).Send(w)
}
