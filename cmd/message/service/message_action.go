package service

import "context"

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) MessageActionService {
	return MessageActionService{ctx: ctx}
}

func (messageActionService MessageActionService) MessageAction() {

}
