package v1

type UserHandler interface {
}

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}
