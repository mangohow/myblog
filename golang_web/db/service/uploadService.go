package service

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"path"
)

type UploadServicer interface {
	UploadImage(string, io.Reader) (string, error)
}

/*
	OssUploadService: 阿里云OSS对象存储服务存储图片
 */
type OssUploadService struct {
	endPoint string
	accessKeyId string
	accessKeySecret string
	bucketName string
	client *oss.Client
}

func NewOssUploadService(endPoint, accessKeyId, accessKeySecret, bucketName string) *OssUploadService {
	o := &OssUploadService{
		endPoint: endPoint,
		accessKeyId: accessKeyId,
		accessKeySecret: accessKeySecret,
		bucketName: bucketName,
	}
	o.init()

	return o
}

func (o *OssUploadService) init() {
	client, err := oss.New(o.endPoint, o.accessKeyId, o.accessKeySecret, oss.UseCname(true))
	if err != nil {
		panic(fmt.Errorf("Create oss client failed:%v", err))
	}
	o.client = client
}

func (o *OssUploadService) UploadImage(objName string, reader io.Reader) (string, error) {
	bucket, err := o.client.Bucket(o.bucketName)
	if err != nil {
		panic(fmt.Errorf("Get bucket failed:%v", err))
	}
	err = bucket.PutObject(objName, reader)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s/%s", o.endPoint, objName), nil
}


/*
	LocalUploadService: 本地存储图片
 */
type LocalUploadService struct {
	filePath string            // ./images
	netBasePath string         // http://localhost:8080/images
}

func NewLocalUploadService(filePath, netBasePath string) *LocalUploadService {
	return &LocalUploadService{
		filePath: filePath,
		netBasePath: netBasePath,
	}
}

func (l *LocalUploadService) UploadImage(objName string, reader io.Reader) (string, error) {
	dst := path.Join(l.filePath, objName)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)

	return  path.Join(l.netBasePath, objName) , err
}