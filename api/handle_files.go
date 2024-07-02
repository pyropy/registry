package api

import (
	"github.com/pyropy/registry/registry"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilesRouteHandler struct {
	registry *registry.Registry
}

func NewFileRouteHandler(reg *registry.Registry) *FilesRouteHandler {
	return &FilesRouteHandler{registry: reg}
}

func (h *FilesRouteHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/files", h.GetFileCID)
	r.POST("/files", h.UploadFile)
}

func (h *FilesRouteHandler) GetFileCID(c *gin.Context) {
	resp := make(map[string]string)
	resp["cid"] = "QmPK1s3pNYLi9ERiq3BDxKa4XosgWwFRQUydHUtz4YgpqB"

	c.JSON(http.StatusOK, resp)
}

func (h *FilesRouteHandler) UploadFile(c *gin.Context) {
	filePath := c.PostForm("filePath")
	file, err := c.FormFile("file")
	if err != nil {
		wrapError(c, err)
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		wrapError(c, err)
		return
	}

	cid, err := h.registry.UploadFile(c, filePath, fileReader)
	if err != nil {
		wrapError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cid": cid,
	})
}

func wrapError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
