package impl

import (
	"context"
	pb "github.com/Katherine-988/subscription_proto/proto"
	"github.com/Katherine-988/tools"
	"time"
)

type Task struct {
	TraceID        string
	OptionType     int32
	UserID         string
	ProductID      int32
	RetryTimes     int
	WriteTimestamp int64
	NeedFeedback   bool
}

func Subscription(c context.Context, req *pb.SubscriptionReq) (*pb.SubscriptionRsp, error) {
	rsp := &pb.SubscriptionRsp{
		Code: pb.RetCode_SUCCESS,
		Msg:  "success",
	}

	task := Task{
		TraceID:        req.Head.TraceID,
		OptionType:     int32(req.OptionType),
		UserID:         req.UserID,
		ProductID:      req.ProductID,
		RetryTimes:     0,
		WriteTimestamp: time.Now().Unix(),
		NeedFeedback:   req.NeedFeedback,
	}
	err := tools.KafkaMgr.Write(c, "task_topic", task)
	if err != nil {
		rsp.Code = pb.RetCode_ERROR
		rsp.Msg = err.Error()
		tools.Errorln(err)
		return rsp, nil
	}
	return rsp, nil
}
