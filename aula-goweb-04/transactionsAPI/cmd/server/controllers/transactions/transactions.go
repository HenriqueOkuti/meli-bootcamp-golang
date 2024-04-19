package controller_transactions

import (
	te "goweb02tarde/internals/repositories/entities"
	ts "goweb02tarde/internals/services/transactions"
	web "goweb02tarde/pkg"
	tv "goweb02tarde/utils/validators"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITransactionsController interface {
	GetAllTransactions(c *gin.Context)
	GetOneTransaction(c *gin.Context)
	InsertTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
	PartialUpdateTransaction(c *gin.Context)
}

type Controller struct {
	transactionsService ts.ITransactionsService
}

func NewTransactionController() ITransactionsController {
	return &Controller{
		transactionsService: ts.NewTransactionsService(),
	}
}

/*
Get all transactions godoc
@Summary Get all transactions
@Description Get all transactions
@Tags transactions
@Accept json
@Produce json
@Success 200 {array} Transaction
@Router /transactions [get]
*/
func (c *Controller) GetAllTransactions(gc *gin.Context) {
	response := web.NewResponse()
	response.JSON(http.StatusOK, c.transactionsService.GetAllTransactions())
	gc.JSON(http.StatusOK, response)
}

func (c *Controller) GetOneTransaction(gc *gin.Context) {
	response := web.NewResponse()

	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	transaction, err := c.transactionsService.GetOneTransaction(id)
	if err != nil {
		response.JSONError(http.StatusNotFound, err)
		gc.JSON(http.StatusNotFound, response)
		return
	}

	response.JSON(http.StatusOK, transaction)
	gc.JSON(http.StatusOK, response)
}

func (c *Controller) InsertTransaction(gc *gin.Context) {
	response := web.NewResponse()

	var transaction te.TransactionDTO

	if err := gc.BindJSON(&transaction); err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	if err := tv.ValidateTransaction(&transaction); err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	result := c.transactionsService.InsertTransaction(transaction)
	response.JSON(http.StatusCreated, result)

	gc.JSON(http.StatusCreated, response)
}

func (c *Controller) UpdateTransaction(gc *gin.Context) {
	response := web.NewResponse()

	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	var transaction te.TransactionDTO
	if err := gc.BindJSON(&transaction); err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	if err := tv.ValidateTransaction(&transaction); err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.transactionsService.UpdateTransaction(id, transaction)
	if err != nil {
		response.JSONError(http.StatusNotFound, err)
		gc.JSON(http.StatusNotFound, response)
		return
	}

	response.JSON(http.StatusOK, result)
	gc.JSON(http.StatusOK, response)
}

func (c *Controller) DeleteTransaction(gc *gin.Context) {
	response := web.NewResponse()

	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.transactionsService.DeleteTransaction(id)
	if err != nil {
		response.JSONError(http.StatusNotFound, err)
		gc.JSON(http.StatusNotFound, response)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (c *Controller) PartialUpdateTransaction(gc *gin.Context) {
	response := web.NewResponse()

	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	var transaction te.PartialTransactionDTO
	if err := gc.BindJSON(&transaction); err != nil {
		response.JSONError(http.StatusBadRequest, err)
		gc.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.transactionsService.PartialUpdateTransaction(id, transaction)
	if err != nil {
		response.JSONError(http.StatusNotFound, err)
		gc.JSON(http.StatusNotFound, response)
		return
	}

	response.JSON(http.StatusOK, result)
	gc.JSON(http.StatusOK, response)
}
