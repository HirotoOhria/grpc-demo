package handler

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pancake.maker/api/gen/api"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type bakerHandler struct {
	report *report
}

type report struct {
	// To Respond multi baking.
	sync.Mutex
	// Map of Menu and baking count.
	data map[api.Pancake_Menu]int
}

func NewBakerHandler() *bakerHandler {
	return &bakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

func (h *bakerHandler) Bake(ctx context.Context, r *api.BakeRequest) (*api.BakeResponse, error) {
	if r.Menu == api.Pancake_UNKNOWN || r.Menu > api.Pancake_BACON_AND_CHEESE {
		return nil, status.Errorf(codes.InvalidArgument, "パンケーキを選択してください")
	}

	now := time.Now()
	h.report.Lock()
	h.report.data[r.Menu] = h.report.data[r.Menu] + 1
	h.report.Unlock()

	res := &api.BakeResponse{
		Pancake: &api.Pancake{
			ChefName:       "ohira",
			Menu:           r.Menu,
			TechnicalScore: rand.Float32(),
			CreateTime: &timestamp.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}

	return res, nil
}

func (h *bakerHandler) Report(ctx context.Context, r *api.ReportRequest) (*api.ReportResponse, error) {
	counts := make([]*api.Report_BakeCount, 0)

	h.report.Lock()
	for menu, count := range h.report.data {
		count := &api.Report_BakeCount{
			Menu:  menu,
			Count: int32(count),
		}

		counts = append(counts, count)
	}
	h.report.Unlock()

	res := &api.ReportResponse{
		Report: &api.Report{
			BakeCounts: counts,
		},
	}

	return res, nil
}
