package controller

import (
	"context"
	"errors"

	pb "github.com/architectural-practices/proto-test/pkg/pb/hello-test"
)

type HelloImpl struct{}

func (s *HelloImpl) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRsp, error) {
	if req.GetOpenid() == "" {
		return nil, errors.New("openid is nil")
	}
	return &pb.HelloRsp{}, nil
}

func (s *HelloImpl) Test(ctx context.Context, req *pb.TestReq) (*pb.TestRsp, error) {
	return nil, nil
}
