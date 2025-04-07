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

// @Summary Update user
// @Description Обновить пользователя
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body usecase.UpdateUserInput true "Update user input"
// @Success 200 {object} nil
// @Failure 400 {object} ErrorResponse
// @Router /user/{id} [put]
func (s *Server) buildUpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc  = usecase.NewUpdateUserUseCase(repo.NewUserDB(s.db), s.ctxTimeout)
			act = action.NewUpdateUserAction(uc, s.log)
		)

		buildParams(c, "id")

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

// @Summary Get user by ID
// @Description Получить пользователя по ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} usecase.FindUserByIDOutput
// @Failure 400 {object} ErrorResponse
// @Router /user/{id} [get]
func (s *Server) buildGetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc  = usecase.NewFindUserByIDUseCase(repo.NewUserDB(s.db), s.ctxTimeout)
			act = action.NewFindUserByIDAction(uc, s.log)
		)

		buildParams(c, "id")

		act.Execute(c.Writer, c.Request)
	}
}

// @Summary Get all users
// @Description Получить всех пользователей
// @Tags user
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param name query string false "Name"
// @Param surname query string false "Surname"
// @Param patronymic query string false "Patronymic"
// @Param age query int false "Age"
// @Param gender query string false "Gender"
// @Param nationality query string false "nationality"
// @Success 200 {object} usecase.FindAllUsersOutput
// @Failure 400 {object} ErrorResponse
// @Router /user [get]
func (s *Server) buildGetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc  = usecase.NewFindAllUsersUseCase(repo.NewUserDB(s.db), s.ctxTimeout)
			act = action.NewFindAllUserAction(uc, s.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}
