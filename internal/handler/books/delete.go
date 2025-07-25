package books

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.srv.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}
