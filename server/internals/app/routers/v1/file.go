package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iads/server/pkg/config"
	"net/http"
	"path"
)

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//fileName := file.Filename
	savePath := path.Join(config.ConfValue.UploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	c.String(http.StatusOK, "上传文件成功")
}

func FileDownload(c *gin.Context) {
	filename := c.Request.PostFormValue("path")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filename)
}
