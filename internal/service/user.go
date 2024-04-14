package service

import (
	"context"
	"rose/internal/dto"
)

func (s *Service) UserDetail(ctx context.Context, uid int64) (*dto.UserDetailResp, error) {

	return &dto.UserDetailResp{
		User: &dto.User{
			UserId:   uid,
			UserName: "hello rose",
		},
	}, nil

}
