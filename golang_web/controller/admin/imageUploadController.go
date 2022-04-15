package admin

import (
	"blog_web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/3 11:08
* @Desc: 图片上传的router
 */

type ImageUploadController struct {
	baseImagePath string
	firstPicPath  string
	blogImgPath   string
	icon          string
}

func NewImageUploadRouter() *ImageUploadController {
	ipr := &ImageUploadController{
		firstPicPath: "/images/firstPic/",
		blogImgPath:  "/images/blogImages/",
		icon: "/images/icons/",
	}
	p := utils.GlobalServerConf.Server.ImagePath
	if(strings.HasSuffix(p, "/")) {
		ipr.baseImagePath = p[:len(p) - 1]
	} else {
		ipr.baseImagePath = p
	}

	return ipr
}

// 上传博客首图
func (ir *ImageUploadController) UploadImage(ctx *gin.Context) {
	_, netpath, err := uploadImage(ctx, ir.baseImagePath, ir.firstPicPath)
	if checkError(err, "Upload image error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 上传博客中的图片
func (ir *ImageUploadController) UploadBlogImage(ctx *gin.Context) {
	_, netpath, err := uploadImage(ctx, ir.baseImagePath, ir.blogImgPath)
	if checkError(err, "Upload blog image error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 上传图标
func (ir *ImageUploadController) UploadIcon(ctx *gin.Context) {
	_, netpath, err := uploadImage(ctx, ir.baseImagePath, ir.icon)
	if checkError(err, "Upload icon error") {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ConvertPathSeparator(netpath))
}

// 接收图片并保存
func uploadImage(ctx *gin.Context, imgPath, imgRelativePath string) (string, string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", "", err
	}

	now := time.Now().Unix()

	fp, suf := utils.FileSuffixSplit(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", fp, now, suf)
	dst := fmt.Sprintf("%s%s%s", imgPath, imgRelativePath, filename)
	utils.Logger().Debug("dest path:%s", dst)
	netPath := fmt.Sprintf("%s%s%s", utils.GlobalServerConf.Server.ImageBaseUrl, imgRelativePath, filename)
	utils.Logger().Debug("net path:%s", netPath)
	utils.Logger().Debug("image base url:%s", imgPath)
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		return "", "", err
	}

	return dst, netPath, nil
}
