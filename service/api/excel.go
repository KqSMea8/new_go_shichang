package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"xinxin/service/common"
	"xinxin/service/excelmgr"
)

func UploadFile(ctx *gin.Context) {
	// 读取Excel
	file, header, err := ctx.Request.FormFile("file")
	filename := header.Filename

	// 保存Excel到上传目录
	path := "./input/" + filename
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, &common.Response{Code: -1, Message: err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, &common.Response{Code: -1, Message: err.Error()})
		return
	}

	// 读取Excel内容
	xlsImpl := new(excelmgr.FileImporter)
	data, err := xlsImpl.LoadData(path)

	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, &common.Response{Code: -1, Message: err.Error()})
	}else{
		ctx.JSON(http.StatusOK, &common.Response{Code: 0, Message: "", Data:data})
	}
}

func DownloadFile(ctx *gin.Context) {
}

func DownloadVendor(ctx *gin.Context) {
}

func UploadVendor(ctx *gin.Context) {
}

func DownloadCost(ctx *gin.Context) {
}

func DownloadPrice(ctx *gin.Context) {
}

func UploadPrice(ctx *gin.Context) {
}