package utils

/*
* @Author: mgh
* @Date: 2022/3/3 11:35
* @Desc:
 */

import (
	"errors"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	Path "path/filepath"
)

// 提取路径中的路径名和后缀 "F:\test\img.jpg" ==> ("F:\test\img", ".jpg")
func FileSuffixSplit(filepath string) (string, string) {
	for i := 0; i < len(filepath); i++ {
		if filepath[i] == '.' {
			return filepath[:i], filepath[i:]
		}
	}

	return filepath, ""
}

// 将 "\" 转为 url中的"/"
func ConvertPathSeparator(path string) string {
	strby := []byte(path)
	for i := 0; i < len(strby); i++ {
		if strby[i] == '\\' {
			strby[i] = '/'
		}
	}

	return string(strby)
}

// 该函数用于对图片进行缩放, 如果图像的宽带大于width，则不需要缩放
// 返回值： string: 缩放后的图像名字，不带路径
func ImageScale(path string, width int) (string, error) {
	pref, suf := FileSuffixSplit(path)
	if suf != ".jpg" && suf != ".jpeg" && suf != ".png" {
		return "", errors.New("Unsupported image type.")
	}

	image, err := imgio.Open(path)
	if err != nil {
		return "", err
	}
	bounds := image.Bounds()
	x := bounds.Dx()
	if x <= width {
		return "", nil
	}

	y := bounds.Dy()
	height := (width * y) / x

	resizedImage := transform.Resize(image, width, height, transform.Linear)

	resizedImgPath := fmt.Sprintf("%s%s%s", pref, "_resized", suf)
	var encoder imgio.Encoder
	if suf == ".jpg" || suf == ".jpeg" {
		encoder = imgio.JPEGEncoder(100)
	} else if suf == ".png" {
		encoder = imgio.PNGEncoder()
	}
	err = imgio.Save(resizedImgPath, resizedImage, encoder)
	if err != nil {
		return "", err
	}

	return Path.Base(resizedImgPath), nil
}
