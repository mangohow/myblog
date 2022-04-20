package admin

import (
	"blog_web/db/service"
	"blog_web/response"
	"blog_web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"path"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/3 11:08
* @Desc: 图片上传的router
 */

type ImageUploadController struct {
	uploadService service.UploadServicer
}

func NewAliOSSImageUploadRouter() *ImageUploadController {
	srv := service.NewOssUploadService(viper.GetString("aliOSS.endPoint"),
		viper.GetString("aliOSS.accessKeyId"),
		viper.GetString("aliOSS.accessKeySecret"),
		viper.GetString("aliOSS.bucket"))
	return &ImageUploadController{
		uploadService: srv,
	}
}

func NewLocalImageUploadRouter() *ImageUploadController {
	srv := service.NewLocalUploadService(viper.GetString("server.imagePath") + "images",
		viper.GetString("server.imageBaseUrl") + "/images")
	return &ImageUploadController{
		uploadService: srv,
	}
}

// 上传博客首图
func (ir *ImageUploadController) UploadImage(ctx *gin.Context) {
	netpath, err := ir.uploadImage(ctx, "firstPic")
	if response.CheckError(err, "Upload image error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 上传博客中的图片
func (ir *ImageUploadController) UploadBlogImage(ctx *gin.Context) {
	netpath, err := ir.uploadImage(ctx, "blogImages")
	if response.CheckError(err, "Upload blog image error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 上传图标
func (ir *ImageUploadController) UploadIcon(ctx *gin.Context) {
	netpath, err := ir.uploadImage(ctx, "icons")
	if response.CheckError(err, "Upload icon error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 接收图片并保存   dir: firstPic, ico
func (ir *ImageUploadController) uploadImage(ctx *gin.Context, dir string) (string, error) {
	file, err := ctx.FormFile("file")     // filename: bg.jpg
	if err != nil {
		return "", err
	}

	utils.Logger().Debug("%v", file.Filename)
	// 在文件名后添加时间戳后缀，防止重复
	now := time.Now().Unix()
	fp, suf := utils.FileSuffixSplit(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", fp, now, suf)   // filename: bg_156435453.jpg
	utils.Logger().Debug("%v", filename)
	objName := path.Join(dir, filename)                    // objName: firstPic/bg_156435453.jpg
	utils.Logger().Debug("objName:%v", objName)
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	maxImageSize := viper.GetInt64("aliOSS.maxImageSize")
	// 如果图片超出了大小，对图片进行一个缩放，缩放为1920*1080
	if( maxImageSize != 0 && file.Size > maxImageSize) {
		buf, err := utils.ImageScale(f, 1920)
		if err != nil {
			return "", err
		}
		return ir.uploadService.UploadImage(objName, buf)
	}

	return ir.uploadService.UploadImage(objName, f)
}
