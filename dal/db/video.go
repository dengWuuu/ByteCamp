/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:27:35
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 22:25:27
 * @FilePath: /ByteCamp/dal/db/video.go
 * @Description: 视频实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"bytes"
	"fmt"
	uuid "github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/plugin/optimisticlock"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	UploadTime    time.Time
	Title         string
	Version       optimisticlock.Version
	FavoriteCount int64
	CommentCount  int64
}

const VideoCount = 10

// GetVideoByIds 根据视频ID获取视频信息
func GetVideoByIds(vids []int64) (resp []*Video, err error) {
	err = DB.Where("ID in ?", vids).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetVideosByLastTime
// 依据一个时间，来获取这个时间之前的一些视频
func GetVideosByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, VideoCount)
	result := DB.Where("upload_time<?", lastTime).Order("upload_time desc").Limit(VideoCount).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}

// GetPublishListByAuthorId 获取用户发布列表
func GetPublishListByAuthorId(authorId int64) ([]Video, error) {
	var data []Video
	//初始化db
	//Init()
	result := DB.Where(&Video{AuthorId: authorId}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

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

// Save 保存视频记录
func Save(video Video) error {
	result := DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSnapshot(videoPath string, frameNum int) (cover io.Reader, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return nil, err
	}
	err = os.RemoveAll(videoPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
