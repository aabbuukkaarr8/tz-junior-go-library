package books

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	booksService "tz-junior-go-library/internal/service/books"
	"tz-junior-go-library/pkg/validator"
)

func (m *CreateRequest) ToSrv() booksService.CreateModel {
	return booksService.CreateModel{
		Title:  m.Title,
		Author: m.Author,
		Pages:  m.Pages,
	}
}
func (m *Model) FromSrv(srv *booksService.Model) {
	m.ID = srv.ID
	m.Title = srv.Title
	m.Author = srv.Author
	m.Pages = srv.Pages
	m.PublishedAt = srv.PublishedAt
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	err := validator.BindJSON(&req, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("[Validator] Invalid JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные входные данные"})
		return
	}
	reqSrv := req.ToSrv()
	createdReqSrv, err := h.srv.Create(reqSrv)
	if err != nil {
		logrus.WithError(err).Error("[Create] cant create service")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdReq := Model{}
	createdReq.FromSrv(createdReqSrv)
	c.JSON(http.StatusCreated, createdReq)

}
