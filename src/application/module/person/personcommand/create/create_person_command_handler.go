package create

import "golang_persons-api/src/domain/person/command"

// NewCreatePersonCommandHandler initializes a new CreatePersonCommandHandler.
func NewCreatePersonCommandHandler( /* service CreatePersonService */ ) CreatePersonCommandHandler {
	return CreatePersonCommandHandler{
		// service: service,
	}
}

type CreatePersonCommandHandler struct {
	service CreatePersonService
}

func (CreatePersonCommandHandler) Handle(command command.Command) error {
	return nil
}

type CreatePersonService struct {
}
