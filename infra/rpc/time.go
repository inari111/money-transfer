package rpc

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TimestampFrom(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   0,
	}
}
