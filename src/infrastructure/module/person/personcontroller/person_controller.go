package personcontroller

import (
	"errors"
	"golang_persons-api/src/application/module/person/commands/create"
	"golang_persons-api/src/application/module/person/commands/delete"
	"golang_persons-api/src/application/module/person/commands/update"
	"golang_persons-api/src/application/module/person/queries/getall"
	"golang_persons-api/src/application/module/person/queries/getone"
	"golang_persons-api/src/domain/module/person"
	"golang_persons-api/src/domain/module/person/command"
	"golang_persons-api/src/domain/module/person/query"
	exception "golang_persons-api/src/infrastructure/http/httperror"
	"golang_persons-api/src/infrastructure/http/httpresponses"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	message = gin.H{
		"test": "ok",
	}
)

// PersonController struct handles person requests
type PersonController struct {
	commandBus command.Bus
	queryBus   query.Bus
}

// NewPersonController creates controller for person requests
func NewPersonController(commandBus command.Bus, queryBus query.Bus) *PersonController {
	return &PersonController{commandBus, queryBus}
}

// GetOnePerson func gets person by id
func (ctrl *PersonController) GetOnePerson(c *gin.Context) {

	aPersonquery := getone.FindPersonQuery{ID: c.Param("person_id")}

	person, err := ctrl.queryBus.Dispatch(c, aPersonquery)
	if err != nil {
		theErr := exception.NewInternalServerError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}
	response := httpresponses.NewHTTPOkResponse("Person found successfully", person)
	c.JSON(response.Status(), response)
}

// GetAllPersons func gets all persons
func (ctrl *PersonController) GetAllPersons(c *gin.Context) {

	allPersonsQuery := getall.FindPersonsQuery{}
	c.JSON(http.StatusOK, allPersonsQuery)
}

// CreatePerson func creates a new person
func (ctrl *PersonController) CreatePerson(c *gin.Context) {

	var createPersonCommand create.CreatePersonCommand
	if err := c.ShouldBindJSON(&createPersonCommand); err != nil {
		theErr := exception.NewUnprocessableEntityError(err.Error())
		// theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	handleResponse(
		c,
		ctrl.commandBus.Dispatch(c, createPersonCommand),
		httpresponses.NewHTTPCreatedResponse("Person created successfully"),
	)
}

// UpdatePerson func updates a person
func (ctrl *PersonController) UpdatePerson(c *gin.Context) {
	var updatePersonCommand update.UpdatePersonCommand
	if err := c.ShouldBindJSON(&updatePersonCommand); err != nil {
		// theErr := exception.NewUnprocessableEntityError(err.Error())
		theErr := exception.NewUnprocessableEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	updatePersonCommand.ID = c.Param("person_id")

	handleResponse(
		c,
		ctrl.commandBus.Dispatch(c, updatePersonCommand),
		httpresponses.NewHTTPCreatedResponse("Person updated successfully"),
	)
}

// DeletePerson func deletes a person
func (ctrl *PersonController) DeletePerson(c *gin.Context) {
	deletePersonCommand := delete.DeletePersonCommand{ID: c.Param("person_id")}

	handleResponse(
		c,
		ctrl.commandBus.Dispatch(c, deletePersonCommand),
		httpresponses.NewHTTPCreatedResponse("Person deleted successfully"),
	)
}

func handleResponse(c *gin.Context, err error, response httpresponses.Message) {

	if err != nil {

		switch {

		case
			errors.Is(err, person.ErrInvalidPersonID),
			errors.Is(err, person.ErrEmptyPersonFirstname),
			errors.Is(err, person.ErrEmptyPersonLastname),
			errors.Is(err, person.ErrInvalidPersonAge):

			theErr := exception.NewBadRequestError(err.Error())
			c.JSON(theErr.Status(), theErr)
			return

		default:

			theErr := exception.NewInternalServerError(err.Error())
			c.JSON(theErr.Status(), theErr)
			return

		}
	}

	c.JSON(response.Status(), response)
}
