package club

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test-api/club/helper"
	"test-api/club/service/club/form"
	"test-api/club/service/club/usecase"
)

type clubHandler struct {
	clubUseCase usecase.ClubUseCase
}

func NewClubHandler(clubUseCase usecase.ClubUseCase) *clubHandler {
	return &clubHandler{clubUseCase}
}

func (h *clubHandler) CreateClubHandler(c *gin.Context) {
	var input form.Club
	err := c.ShouldBind(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.clubUseCase.CreateClubUseCase(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), result))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *clubHandler) UpdateRecordMatchHandler(c *gin.Context) {
	var input form.RecordMatch
	err := c.ShouldBind(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := h.clubUseCase.UpdateRecordMatchUserCase(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)

}

func (h *clubHandler) GetAllClubHandler(c *gin.Context) {
	result, err := h.clubUseCase.GetAllClubUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GetResponse(http.StatusInternalServerError, false, err.Error(), false))
		return
	}

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}

func (h *clubHandler) ContainLetterHandler(c *gin.Context) {
	var input form.ContainLetter
	err := c.ShouldBind(&input)
	if err != nil {
		errorMessages := helper.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := helper.GetResponse(http.StatusBadRequest, false, errorMessage, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := h.clubUseCase.ContainLetterUseCase(input)

	res := helper.GetSuccessResponse(result)
	c.JSON(http.StatusOK, res)
}
