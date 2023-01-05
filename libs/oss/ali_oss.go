package oss

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/url"
	"path/filepath"
	"strconv"
)

var NotAbsPath = errors.New("文件路径不是绝对地址")

type AliOssApi struct {
	EndPoint        string // 阿里oss内网节点地址
	PublicEndPoint  string // 阿里外网节点地址
	AccessKeyId     string // 访问密钥id
	AccessKeySecret string // 访问密钥
	BucketName      string // oss 包名称
	Host            string // 拼接的外网可访问的地址
	bucket          *oss.Bucket
}

// NewAliOssApi 阿里oss的构造方法 创建新的oss对象
func NewAliOssApi(endPoint, publicEndPoint, accessKeyId, accessKeySecret, bucketName string) (*AliOssApi, error) {
	ossClient, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := ossClient.Bucket("test")
	if err != nil {
		return nil, err
	}
	host := "https://" + bucketName + "." + publicEndPoint + "/"
	return &AliOssApi{
		EndPoint: endPoint, PublicEndPoint: publicEndPoint, AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BucketName: bucketName, bucket: bucket, Host: host}, nil
}

// GetFullURL 获取ossKey的外网可访问的完整地址
func (a AliOssApi) GetFullURL(ossKey string) string {
	return a.Host + url.QueryEscape(ossKey)
}

// TranceOssKey 将url地址转换为ossKey
func (a AliOssApi) TranceOssKey(fileUrl string) (string, error) {
	uri, err := url.ParseRequestURI(fileUrl)
	if err != nil {
		return "", err
	}
	return string([]byte(uri.Path)[1:]), nil
}

// UploadLocalFile 上传本地文件
// filePath 为本地文件完整路径
func (a AliOssApi) UploadLocalFile(ossKey string, filePath string) error {
	if !filepath.IsAbs(filePath) {
		return NotAbsPath
	}
	return a.bucket.PutObjectFromFile(ossKey, filePath)
}

// UploadFileObj 上传文件对象
func (a AliOssApi) UploadFileObj(ossKey string, fileObj io.Reader) error {
	return a.bucket.PutObject(ossKey, fileObj)
}

// AppendFileObj 追加上传文件
func (a AliOssApi) AppendFileObj(ossKey string, fileObj io.Reader, appendPosition int64) (int64, error) {
	return a.bucket.AppendObject(ossKey, fileObj, appendPosition)
}

// GetAppendPosition 获取追加位置
func (a AliOssApi) GetAppendPosition(ossKey string) (int64, error) {
	meta, err := a.bucket.GetObjectDetailedMeta(ossKey)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(meta.Get("X-Oss-Next-Append-Position"), 10, 64)
}

// DownloadFileObj 下载文件到缓存
func (a AliOssApi) DownloadFileObj(ossKey string) (*bytes.Buffer, error) {
	object, err := a.bucket.GetObject(ossKey)
	if err != nil {
		return nil, err
	}
	defer object.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, object)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// DownloadFileLocal 下载文件到本地
func (a AliOssApi) DownloadFileLocal(ossKey, filePath string) error {
	if !filepath.IsAbs(filePath) {
		return NotAbsPath
	}
	return a.bucket.GetObjectToFile(ossKey, filePath)
}

// ExistFile 判断文件是否存在
func (a AliOssApi) ExistFile(ossKey string) (bool, error) {
	return a.bucket.IsObjectExist(ossKey)
}

// DeleteFile 删除单个文件
func (a AliOssApi) DeleteFile(ossKey string) error {
	return a.bucket.DeleteObject(ossKey)
}
