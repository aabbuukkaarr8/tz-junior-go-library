package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tz-junior-go-library/pkg/validator"
)

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[Get By id]Book Not Found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var input UpdateBook
	err = validator.BindJSON(&input, c.Request)
	if err != nil {
		logrus.WithError(err).Errorf("[Validate] bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSrv, err := h.srv.Update(input.UpdateSrv(id))
	if err != nil {
		logrus.WithError(err).Errorf("[Update ] Error updating product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := Model{}
	result.FromSrv(updatedSrv)
	c.JSON(http.StatusOK, result)
}
