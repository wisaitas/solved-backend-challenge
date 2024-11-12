package three

import (
	"context"
	"encoding/json"

	pb "github.com/wisaitas/solved-backend-challenge/three/proto"
)

type BeefServer struct {
	pb.UnimplementedBeefCounterServer
}

func (s *BeefServer) CountBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	jsonResult := PireFireDire(req.Text)

	var result map[string]interface{}
	json.Unmarshal([]byte(jsonResult), &result)

	beefCount := make(map[string]int32)
	if beefMap, ok := result["beef"].(map[string]interface{}); ok {
		for k, v := range beefMap {
			if count, ok := v.(float64); ok {
				beefCount[k] = int32(count)
			}
		}
	}

	return &pb.BeefResponse{
		BeefCount: beefCount,
	}, nil
}
