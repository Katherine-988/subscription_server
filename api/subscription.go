package api

import (
	"context"
	pb "github.com/Katherine-988/subscription_proto/proto"
	"github.com/Katherine-988/subscription_server/impl"
)

type SubscriptionService struct {
	pb.SubscriptionServiceServer
}

func (m *SubscriptionService) Subscription(c context.Context, req *pb.SubscriptionReq) (*pb.SubscriptionRsp, error) {
	return impl.Subscription(c, req)
}
