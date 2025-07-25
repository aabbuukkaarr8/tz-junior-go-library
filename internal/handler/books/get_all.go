package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAll(c *gin.Context) {
	books, err := h.srv.GetAll()
	if err != nil {
		logrus.WithError(err).Errorf("[Get All]Error getting books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []Model
	for book := range books {
		var b Model
		b.FromSrv(&books[book])
		response = append(response, b)
	}
	c.JSON(http.StatusOK, response)
}
