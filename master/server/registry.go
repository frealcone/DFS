package master_server

import (
	"container/heap"
	"fmt"
	"time"
)

type ChunkServer struct {
	address  string
	port     int
	weight   uint32
	startAt  uint64
	workload uint32
}

// Registry manages infomation about running chunk servers.
type Registry interface {
	Register(ChunkServer) error
	Next() *ChunkServer
	All() []ChunkServer
}

// LocalRegistry stores all chunk servers' infomation in master server's
// run time memory space.
type LocalRegistry []ChunkServer

func (lr *LocalRegistry) Register(cs ChunkServer) error {
	for _, s := range *lr {
		if s.address == cs.address && s.port == cs.port {
			return nil
		}
	}

	heap.Push(lr, cs)
	return nil
}

func (lr *LocalRegistry) Next() *ChunkServer {
	cs := heap.Pop(lr).(ChunkServer)
	cs.workload++
	heap.Push(lr, cs)
	return &cs
}

func (lr *LocalRegistry) All() []ChunkServer {
	r := make([]ChunkServer, len(*lr))
	copy(r, *lr)
	return r
}

func (lr LocalRegistry) Len() int {
	return len(lr)
}

func (lr LocalRegistry) Swap(i, j int) {
	lr[i], lr[j] = lr[j], lr[i]
}

func (lr LocalRegistry) Less(i, j int) bool {
	return lr[i].Compare(lr[j]) < 0
}

func (lr *LocalRegistry) Push(x interface{}) {
	if cs, ok := x.(ChunkServer); ok {
		*lr = append(*lr, cs)
	}
}

func (lr *LocalRegistry) Pop() interface{} {
	n := len(*lr)
	r := (*lr)[n-1]
	*lr = (*lr)[:n-1]
	return r
}

func (cs ChunkServer) Compare(server ChunkServer) int {
	if cs.workload/cs.weight != server.workload/server.weight {
		return int(cs.workload/cs.weight) - int(server.workload/server.weight)
	}
	if cs.weight != server.weight {
		return int(server.weight) - int(cs.weight)
	}
	if cs.startAt < server.startAt {
		return -1
	} else if cs.startAt > server.startAt {
		return 1
	}
	return 0
}

func NewChunkServer(address string, port int, weight uint32) ChunkServer {
	return ChunkServer{
		address:  address,
		port:     port,
		weight:   weight,
		startAt:  uint64(time.Now().UnixNano()),
		workload: 0,
	}
}

func (cs ChunkServer) GetIPAddress() string {
	return fmt.Sprintf("%s:%d", cs.address, cs.port)
}
