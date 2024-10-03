package master_server

import (
	"context"
	"fmt"

	"github.com/frealcone/DFS/pb"
)

type MasterServer struct {
	pb.UnimplementedMasterServer
	registry Registry
}

func (s *MasterServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	// get all other servers
	chunkServers := s.registry.All()
	chunkServerAddrs := make([]string, len(chunkServers))
	for i := range chunkServers {
		chunkServerAddrs[i] = fmt.Sprintf("%s:%d", chunkServers[i].address, chunkServers[i].port)
	}

	// register current server into registry
	err := s.registry.Register(NewChunkServer(
		req.GetAddress(),
		int(req.GetPort()),
		req.GetWeight(),
	))

	// return results
	return &pb.RegisterResp{
		Addresses: chunkServerAddrs,
	}, err
}
