package action

import (
	"encoding/json"
	"net/http"

	"test-task-user/internal/adapters/api/logging"
	"test-task-user/internal/adapters/api/response"
	"test-task-user/internal/usecase"

	"github.com/sirupsen/logrus"
)

type CreateUserAction struct {
	uc  usecase.CreateUserUseCase
	log *logrus.Logger
}

func NewCreateUserAction(uc usecase.CreateUserUseCase, log *logrus.Logger) CreateUserAction {
	return CreateUserAction{
		uc:  uc,
		log: log,
	}
}

func (a CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_user"
	var input usecase.CreateUserInput

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

	err := a.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewErrorWithErrorStatus(err, w, a.log, logKey, "error when create_user")
		return
	}

	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success create_user")

	response.NewSuccess(nil, http.StatusOK).Send(w)
}
