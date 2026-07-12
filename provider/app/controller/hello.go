package controller

import (
	"context"
	"errors"

	pb "github.com/architectural-practices/proto-test/pkg/pb/hello-test"
	"trpc.group/trpc-go/trpc-go/log"

	"polarisMesh/app/trace"
)

type HelloImpl struct{}

func (s *HelloImpl) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRsp, error) {
	log.Infof("received hello request trace_id=%s", trace.ID(ctx))
	if req.GetOpenid() == "" {
		return nil, errors.New("openid is nil")
	}
	return &pb.HelloRsp{}, nil
}

func (s *HelloImpl) Test(ctx context.Context, req *pb.TestReq) (*pb.TestRsp, error) {
	log.Infof("received test request trace_id=%s", trace.ID(ctx))
	return nil, nil
}
