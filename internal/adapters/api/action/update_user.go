package action

import (
	"encoding/json"
	"net/http"
	"strconv"

	"test-task-user/internal/adapters/api/logging"
	"test-task-user/internal/adapters/api/response"
	"test-task-user/internal/usecase"

	"github.com/sirupsen/logrus"
)

type UpdateUserAction struct {
	uc  usecase.UpdateUserUseCase
	log *logrus.Logger
}

func NewUpdateUserAction(uc usecase.UpdateUserUseCase, log *logrus.Logger) UpdateUserAction {
	return UpdateUserAction{
		uc:  uc,
		log: log,
	}
}

func (a UpdateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "update_user"
	var input usecase.UpdateUserInput

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

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	err = a.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewErrorWithErrorStatus(err, w, a.log, logKey, "error when update_user")
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success update_user")

	response.NewSuccess(nil, http.StatusOK).Send(w)
}
