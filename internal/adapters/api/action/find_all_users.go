package action

import (
	"fmt"
	"net/http"

	"test-task-user/internal/adapters/api/logging"
	"test-task-user/internal/adapters/api/response"
	"test-task-user/internal/usecase"

	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type FindAllUsersAction struct {
	uc  usecase.FindAllUsersUseCase
	log *logrus.Logger
}

func NewFindAllUserAction(uc usecase.FindAllUsersUseCase, log *logrus.Logger) FindAllUsersAction {
	return FindAllUsersAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAllUsersAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_users"
	var input usecase.FindAllUsersInput

	if err := binding.Default(r.Method, binding.MIMEPOSTForm).Bind(r, &input); err != nil {
		response.NewErrorWithErrorStatus(err, w, a.log, logKey, fmt.Sprintf("error when %s", logKey))
		return
	}

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewErrorWithErrorStatus(err, w, a.log, logKey, "error when find_all_users")
		return
	}

	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success find_all_users")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
