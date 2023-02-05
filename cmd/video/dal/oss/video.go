package oss

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func UploadVideo(video *[]byte) (playUrl string, coverUrl string, err error) {
	videoName := uuid.NewV4().String() + ".mp4"
	imageName := uuid.NewV4().String() + ".jpeg"

	err = os.WriteFile(videoName, *video, 0666)
	if err != nil {
		return "", "", err
	}

	imageData, _ := GetSnapshot(videoName, 1)
	if err != nil {
		return "", "", err
	}
	err = VideoBucket.PutObject(videoName, bytes.NewReader(*video))
	if err != nil {
		return "", "", err
	}

	err = ImageBucket.PutObject(imageName, imageData)
	if err != nil {
		return "", "", err
	}

	playUrl = VideoBucketLinkPrefix + videoName
	coverUrl = ImageBucketLinkPrefix + imageName
	return playUrl, coverUrl, nil
}

func GetSnapshot(videoPath string, frameNum int) (cover io.Reader, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("Extract Frame Failed", err)
		return nil, err
	}
	err = os.RemoveAll(videoPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
