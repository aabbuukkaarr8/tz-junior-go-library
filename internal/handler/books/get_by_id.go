package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) GetByTitle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	bs, err := h.srv.GetById(id)
	if err != nil {
		logrus.WithError(err).Errorf("[Get By id]Book Not Found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	b := Model{}
	b.FromSrv(bs)
	c.JSON(http.StatusOK, b)
}
