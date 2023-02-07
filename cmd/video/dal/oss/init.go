package oss

import (
	"fmt"

	"douyin/cmd/video/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	VideoBucket           *oss.Bucket
	ImageBucket           *oss.Bucket
	VideoBucketLinkPrefix string
	ImageBucketLinkPrefix string
)

func Init() {
	ossClient, err := oss.New(config.OSSEndPoint, config.OSSAK, config.OSSSK)
	if err != nil {
		hlog.Fatalf("OSS Init Failed")
		panic(err)
	}

	VideoBucket, err = ossClient.Bucket(config.OSSVideoBucket)
	if err != nil {
		hlog.Fatalf("VideoBucket Init Failed")
		panic(err)
	}
	VideoBucketLinkPrefix = fmt.Sprintf(
		"https://%s.%s/", config.OSSVideoBucket, config.OSSEndPoint)

	ImageBucket, err = ossClient.Bucket(config.OSSImageBucket)
	if err != nil {
		hlog.Fatalf("ImageBucket Init Failed")
		panic(err)
	}
	ImageBucketLinkPrefix = fmt.Sprintf(
		"https://%s.%s/", config.OSSImageBucket, config.OSSEndPoint)
}
