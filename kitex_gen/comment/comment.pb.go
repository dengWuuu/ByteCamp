// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: idl/comment.proto

package comment

import (
	context "context"
	user "douyin/kitex_gen/user"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DouyinCommentActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                     // 用户id
	Token       string  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                                      // 用户鉴权token
	VideoId     int64   `protobuf:"varint,3,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`                  // 视频id
	ActionType  int32   `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`         // 1-发布评论，2-删除评论
	CommentText *string `protobuf:"bytes,5,opt,name=comment_text,json=commentText,proto3,oneof" json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `protobuf:"varint,6,opt,name=comment_id,json=commentId,proto3,oneof" json:"comment_id,omitempty"`      // 要删除的评论id，在action_type=2的时候使用
}

func (x *DouyinCommentActionRequest) Reset() {
	*x = DouyinCommentActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentActionRequest) ProtoMessage() {}

func (x *DouyinCommentActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinCommentActionRequest) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinCommentActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinCommentActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinCommentActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *DouyinCommentActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *DouyinCommentActionRequest) GetCommentText() string {
	if x != nil && x.CommentText != nil {
		return *x.CommentText
	}
	return ""
}

func (x *DouyinCommentActionRequest) GetCommentId() int64 {
	if x != nil && x.CommentId != nil {
		return *x.CommentId
	}
	return 0
}

type DouyinCommentActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	Comment    *Comment `protobuf:"bytes,3,opt,name=comment,proto3,oneof" json:"comment,omitempty"`                      // 评论成功返回评论内容，不需要重新拉取整个列表
}

func (x *DouyinCommentActionResponse) Reset() {
	*x = DouyinCommentActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentActionResponse) ProtoMessage() {}

func (x *DouyinCommentActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinCommentActionResponse) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinCommentActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinCommentActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinCommentActionResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type DouyinCommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                     // 用户鉴权token
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *DouyinCommentListRequest) Reset() {
	*x = DouyinCommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentListRequest) ProtoMessage() {}

func (x *DouyinCommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentListRequest.ProtoReflect.Descriptor instead.
func (*DouyinCommentListRequest) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinCommentListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinCommentListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type DouyinCommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   *string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	CommentList []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` // 评论列表
}

func (x *DouyinCommentListResponse) Reset() {
	*x = DouyinCommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentListResponse) ProtoMessage() {}

func (x *DouyinCommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentListResponse.ProtoReflect.Descriptor instead.
func (*DouyinCommentListResponse) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinCommentListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinCommentListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinCommentListResponse) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  // 视频评论id
	User       *user.User `protobuf:"bytes,2,opt,name=userHandler,proto3" json:"userHandler,omitempty"`                 // 评论用户信息
	Content    string     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                         // 评论内容
	CreateDate string     `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{4}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

var File_idl_comment_proto protoreflect.FileDescriptor

var file_idl_comment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x69, 0x64,
	0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a,
	0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x1b, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x1c, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x33, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x22, 0x74, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x32, 0xca, 0x01, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x72, 0x76, 0x12, 0x60, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_comment_proto_rawDescOnce sync.Once
	file_idl_comment_proto_rawDescData = file_idl_comment_proto_rawDesc
)

func file_idl_comment_proto_rawDescGZIP() []byte {
	file_idl_comment_proto_rawDescOnce.Do(func() {
		file_idl_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_comment_proto_rawDescData)
	})
	return file_idl_comment_proto_rawDescData
}

var file_idl_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_idl_comment_proto_goTypes = []interface{}{
	(*DouyinCommentActionRequest)(nil),  // 0: comment.douyin_comment_action_request
	(*DouyinCommentActionResponse)(nil), // 1: comment.douyin_comment_action_response
	(*DouyinCommentListRequest)(nil),    // 2: comment.douyin_comment_list_request
	(*DouyinCommentListResponse)(nil),   // 3: comment.douyin_comment_list_response
	(*Comment)(nil),                     // 4: comment.Comment
	(*user.User)(nil),                   // 5: userHandler.User
}
var file_idl_comment_proto_depIdxs = []int32{
	4, // 0: comment.douyin_comment_action_response.comment:type_name -> comment.Comment
	4, // 1: comment.douyin_comment_list_response.comment_list:type_name -> comment.Comment
	5, // 2: comment.Comment.userHandler:type_name -> userHandler.User
	0, // 3: comment.CommentSrv.CommentAction:input_type -> comment.douyin_comment_action_request
	2, // 4: comment.CommentSrv.CommentList:input_type -> comment.douyin_comment_list_request
	1, // 5: comment.CommentSrv.CommentAction:output_type -> comment.douyin_comment_action_response
	3, // 6: comment.CommentSrv.CommentList:output_type -> comment.douyin_comment_list_response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_idl_comment_proto_init() }
func file_idl_comment_proto_init() {
	if File_idl_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_idl_comment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_idl_comment_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_idl_comment_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_comment_proto_goTypes,
		DependencyIndexes: file_idl_comment_proto_depIdxs,
		MessageInfos:      file_idl_comment_proto_msgTypes,
	}.Build()
	File_idl_comment_proto = out.File
	file_idl_comment_proto_rawDesc = nil
	file_idl_comment_proto_goTypes = nil
	file_idl_comment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type CommentSrv interface {
	CommentAction(ctx context.Context, req *DouyinCommentActionRequest) (res *DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, req *DouyinCommentListRequest) (res *DouyinCommentListResponse, err error)
}
