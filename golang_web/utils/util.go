package utils

/*
* @Author: mgh
* @Date: 2022/3/3 11:35
* @Desc:
 */

import (
	"bytes"
	"errors"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
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


// ImageScale: 对图片进行缩放
func ImageScale(reader io.Reader, width int) (io.Reader, error) {
	img, suf, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	Logger().Debug("image type:%s", suf)
	if suf != "jpg" && suf != "jpeg" && suf != "png" {
		return nil, errors.New("Unsupported image type.")
	}

	bounds := img.Bounds()
	x := bounds.Dx()
	if x <= width {
		return nil, nil
	}

	y := bounds.Dy()
	height := (width * y) / x

	resizedImage := transform.Resize(img, width, height, transform.Linear)
	Logger().Debug("%v", resizedImage.Bounds().Size())
	buf := bytes.NewBuffer(nil)
	var encoder imgio.Encoder
	if suf == "jpg" || suf == "jpeg" {
		encoder = imgio.JPEGEncoder(100)
	} else if suf == "png" {
		encoder = imgio.PNGEncoder()
	}
	err = encoder(buf, resizedImage)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
