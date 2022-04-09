package utils

import (
	"fmt"
	"testing"
	"time"
)

/*
* @Author: mgh
* @Date: 2022/3/2 15:03
* @Desc:
 */

func TestTokenGen(t *testing.T) {
	token, err := CreateToken(1, "1158446387", time.Second*1)
	fmt.Println(token)
	if err != nil {
		t.Error("createToken:", err)
	}

	time.Sleep(time.Second * 2)
	username, id, ok := VerifyToken(token)
	fmt.Println(username, id, ok)
}

func TestVerify(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IjExNTg0NDYzODciLCJVc2VySWQiOjEsImV4cCI6MTY0NjIwNDkwNCwiaWF0IjoxNjQ2MjA0ODQ0LCJpc3MiOiJtZ2giLCJzdWIiOiJVc2VyX1Rva2VuIn0.kH7WTTl-i39bECFlJb5pemiSq9wJnITg6o9ydJmySBE"
	username, id, ok := VerifyToken(token)
	fmt.Println(username, id, ok)
}

func TestGetFilenameWithoutSuffix(t *testing.T) {
	fp, suffix := FileSuffixSplit("abc.jpg")
	fmt.Println(fp, suffix)
}

func TestImageScale(t *testing.T) {
	path, err := ImageScale("F:\\go_projects\\vue_blog_web\\images\\firstPic\\29.png", 800)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(path)
}
