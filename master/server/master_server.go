package master_server

import (
	"context"
	"fmt"
	"io"

	"github.com/frealcone/DFS/pb"
)

type MasterServer struct {
	pb.UnimplementedMasterServer
	registry Registry
	fs       FileSystem
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

func (s *MasterServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	if err := s.fs.Touch(req.Filename); err != nil {
		return nil, err
	}
	return &pb.CreateResp{}, nil
}

func (s *MasterServer) Discover(stream pb.Master_DiscoverServer) error {
	for {
		entryReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		h, err := s.fs.GetHandle(entryReq.GetFilename(), entryReq.GetOffset())
		if err != nil {
			return err
		}

		entries := []*pb.Entry{}
		for serv, v := range h.chunkServers {
			if h.primary == serv {
				continue
			}

			entries = append(entries, &pb.Entry{
				ChunkName:   h.chunkName,
				ChunkServer: serv,
				Version:     v.Load(),
			})
		}

		stream.Send(&pb.EntryResp{
			Entries: entries,
		})
	}

	return nil
}
