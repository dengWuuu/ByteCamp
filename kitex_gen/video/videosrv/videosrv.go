// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videosrv

import (
	"context"
	video "douyin/kitex_gen/video"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoSrvServiceInfo
}

var videoSrvServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoSrv"
	handlerType := (*video.VideoSrv)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newPublishActionArgs, newPublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
		"GetUserFeed":   kitex.NewMethodInfo(getUserFeedHandler, newGetUserFeedArgs, newGetUserFeedResult, false),
		"GetVideoById":  kitex.NewMethodInfo(getVideoByIdHandler, newGetVideoByIdArgs, newGetVideoByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DouyinPublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoSrv).PublishAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishActionArgs:
		success, err := handler.(video.VideoSrv).PublishAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishActionResult)
		realResult.Success = success
	}
	return nil
}
func newPublishActionArgs() interface{} {
	return &PublishActionArgs{}
}

func newPublishActionResult() interface{} {
	return &PublishActionResult{}
}

type PublishActionArgs struct {
	Req *video.DouyinPublishActionRequest
}

func (p *PublishActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DouyinPublishActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishActionArgs) Unmarshal(in []byte) error {
	msg := new(video.DouyinPublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishActionArgs_Req_DEFAULT *video.DouyinPublishActionRequest

func (p *PublishActionArgs) GetReq() *video.DouyinPublishActionRequest {
	if !p.IsSetReq() {
		return PublishActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishActionResult struct {
	Success *video.DouyinPublishActionResponse
}

var PublishActionResult_Success_DEFAULT *video.DouyinPublishActionResponse

func (p *PublishActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DouyinPublishActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishActionResult) Unmarshal(in []byte) error {
	msg := new(video.DouyinPublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishActionResult) GetSuccess() *video.DouyinPublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DouyinPublishActionResponse)
}

func (p *PublishActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DouyinPublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoSrv).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(video.VideoSrv).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *video.DouyinPublishListRequest
}

func (p *PublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DouyinPublishListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(video.DouyinPublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *video.DouyinPublishListRequest

func (p *PublishListArgs) GetReq() *video.DouyinPublishListRequest {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *video.DouyinPublishListResponse
}

var PublishListResult_Success_DEFAULT *video.DouyinPublishListResponse

func (p *PublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DouyinPublishListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(video.DouyinPublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *video.DouyinPublishListResponse {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DouyinPublishListResponse)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DouyinFeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoSrv).GetUserFeed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserFeedArgs:
		success, err := handler.(video.VideoSrv).GetUserFeed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserFeedResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserFeedArgs() interface{} {
	return &GetUserFeedArgs{}
}

func newGetUserFeedResult() interface{} {
	return &GetUserFeedResult{}
}

type GetUserFeedArgs struct {
	Req *video.DouyinFeedRequest
}

func (p *GetUserFeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DouyinFeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserFeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserFeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserFeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserFeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserFeedArgs) Unmarshal(in []byte) error {
	msg := new(video.DouyinFeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserFeedArgs_Req_DEFAULT *video.DouyinFeedRequest

func (p *GetUserFeedArgs) GetReq() *video.DouyinFeedRequest {
	if !p.IsSetReq() {
		return GetUserFeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserFeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserFeedResult struct {
	Success *video.DouyinFeedResponse
}

var GetUserFeedResult_Success_DEFAULT *video.DouyinFeedResponse

func (p *GetUserFeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DouyinFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserFeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserFeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserFeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserFeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserFeedResult) Unmarshal(in []byte) error {
	msg := new(video.DouyinFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserFeedResult) GetSuccess() *video.DouyinFeedResponse {
	if !p.IsSetSuccess() {
		return GetUserFeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DouyinFeedResponse)
}

func (p *GetUserFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getVideoByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.VideoIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoSrv).GetVideoById(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideoByIdArgs:
		success, err := handler.(video.VideoSrv).GetVideoById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideoByIdResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideoByIdArgs() interface{} {
	return &GetVideoByIdArgs{}
}

func newGetVideoByIdResult() interface{} {
	return &GetVideoByIdResult{}
}

type GetVideoByIdArgs struct {
	Req *video.VideoIdRequest
}

func (p *GetVideoByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.VideoIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetVideoByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetVideoByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetVideoByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoByIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideoByIdArgs) Unmarshal(in []byte) error {
	msg := new(video.VideoIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideoByIdArgs_Req_DEFAULT *video.VideoIdRequest

func (p *GetVideoByIdArgs) GetReq() *video.VideoIdRequest {
	if !p.IsSetReq() {
		return GetVideoByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideoByIdResult struct {
	Success *video.Video
}

var GetVideoByIdResult_Success_DEFAULT *video.Video

func (p *GetVideoByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.Video)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetVideoByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetVideoByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetVideoByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoByIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideoByIdResult) Unmarshal(in []byte) error {
	msg := new(video.Video)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoByIdResult) GetSuccess() *video.Video {
	if !p.IsSetSuccess() {
		return GetVideoByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.Video)
}

func (p *GetVideoByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, Req *video.DouyinPublishActionRequest) (r *video.DouyinPublishActionResponse, err error) {
	var _args PublishActionArgs
	_args.Req = Req
	var _result PublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *video.DouyinPublishListRequest) (r *video.DouyinPublishListResponse, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFeed(ctx context.Context, Req *video.DouyinFeedRequest) (r *video.DouyinFeedResponse, err error) {
	var _args GetUserFeedArgs
	_args.Req = Req
	var _result GetUserFeedResult
	if err = p.c.Call(ctx, "GetUserFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoById(ctx context.Context, Req *video.VideoIdRequest) (r *video.Video, err error) {
	var _args GetVideoByIdArgs
	_args.Req = Req
	var _result GetVideoByIdResult
	if err = p.c.Call(ctx, "GetVideoById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}