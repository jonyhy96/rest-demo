package controllers

import (
	"net/http"
	"rest-demo/models"
)

// ErrorController handles normal error
type ErrorController struct {
	baseController
}

// Error404 handles 404 error
func (c *ErrorController) Error404() {
	e := models.HTTPError{
		Status:  http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	}
	c.returnJSON(http.StatusNotFound, e)
}

// Error500 handles 500 error
func (c *ErrorController) Error500() {
	c.Ctx.Output.SetStatus(http.StatusInternalServerError)
	e := models.HTTPError{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
	c.returnJSON(http.StatusInternalServerError, e)
}
