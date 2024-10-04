package master_client

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/frealcone/DFS/pb"
)

func Entries(client pb.MasterClient, filenames []string, offsets []uint64) ([]*pb.Entry, error) {
	if len(filenames) != len(offsets) {
		return []*pb.Entry{}, fmt.Errorf("number of filenames don't match with number of offsets")
	}
	if len(filenames) <= 0 {
		return []*pb.Entry{}, nil
	}

	entries := []*pb.Entry{}

	stream, err := client.Discover(context.Background())
	if err != nil {
		return entries, err
	}

	var remoteErr error
	// create goroutine waiting for master server's response
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			resp, e := stream.Recv()
			if e != nil {
				if e != io.EOF {
					remoteErr = e
				}
				break
			}
			entries = append(entries, resp.Entries...)
		}
	}()

	// send entry requests to master server
	for i := range filenames {
		err = stream.Send(&pb.EntryReq{
			Filename: filenames[i],
			Offset:   offsets[i],
		})
		if err != nil {
			return []*pb.Entry{}, err
		}
	}

	stream.CloseSend()
	wg.Wait()

	return entries, remoteErr
}

func Touch(client pb.MasterClient, filename string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Create(ctx, &pb.CreateReq{Filename: filename})
	return err
}

func Register(client pb.MasterClient, address string, port int, weight int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Register(ctx, &pb.RegisterReq{
		Address: address,
		Port:    int32(port),
		Weight:  uint32(weight),
	})
	if err != nil {
		return []string{}, err
	}

	return resp.GetAddresses(), nil
}
