package controller

import (
	te "goweb02tarde/internals/repositories/entities"
	ts "goweb02tarde/internals/services"
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

func (c *Controller) GetAllTransactions(gc *gin.Context) {
	gc.JSON(http.StatusOK, c.transactionsService.GetAllTransactions())
}

func (c *Controller) GetOneTransaction(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := c.transactionsService.GetOneTransaction(id)
	if err != nil {
		gc.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, transaction)
}

func (c *Controller) InsertTransaction(gc *gin.Context) {
	var transaction te.TransactionDTO
	if err := gc.BindJSON(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tv.ValidateTransaction(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.transactionsService.InsertTransaction(transaction)

	gc.JSON(http.StatusCreated, result)
}

func (c *Controller) UpdateTransaction(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transaction te.TransactionDTO
	if err := gc.BindJSON(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tv.ValidateTransaction(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.transactionsService.UpdateTransaction(id, transaction)
	if err != nil {
		gc.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (c *Controller) DeleteTransaction(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.transactionsService.DeleteTransaction(id)
	if err != nil {
		gc.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (c *Controller) PartialUpdateTransaction(gc *gin.Context) {
	id, err := strconv.Atoi(gc.Param("id"))
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transaction te.PartialTransactionDTO
	if err := gc.BindJSON(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.transactionsService.PartialUpdateTransaction(id, transaction)
	if err != nil {
		gc.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, result)
}
