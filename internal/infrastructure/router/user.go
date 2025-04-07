package server

import (
	"test-task-user/internal/adapters/api/action"
	"test-task-user/internal/adapters/repo"
	"test-task-user/internal/usecase"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Создать пользователя
// @Tags user
// @Accept json
// @Produce json
// @Param input body usecase.CreateUserInput true "Create user input"
// @Success 200 {object} nil
// @Failure 400 {object} ErrorResponse
// @Router /user [post]
func (s *Server) buildCreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewCreateUserUseCase(
				repo.NewUserDB(s.db),
				s.enrichment,
				s.ctxTimeout,
			)
			act = action.NewCreateUserAction(uc, s.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

// @Summary Delete user
// @Description Удалить пользователя
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} nil
// @Failure 400 {object} ErrorResponse
// @Router /user/{id} [delete]
func (s *Server) buildDeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc  = usecase.NewDeleteUserUseCase(repo.NewUserDB(s.db), s.ctxTimeout)
			act = action.NewDeleteUserAction(uc, s.log)
		)

		buildParams(c, "id")

		act.Execute(c.Writer, c.Request)
	}
}
