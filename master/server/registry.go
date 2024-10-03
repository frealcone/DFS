package master_server

import "time"

type ChunkServer struct {
	address string
	port    int
	weight  uint32
	startAt uint64
}

// Registry manages infomation about running chunk servers.
type Registry interface {
	Register(ChunkServer) error
	Next() *ChunkServer
}

// LocalRegistry stores all chunk servers' infomation in master server's
// run time memory space.
type LocalRegistry []ChunkServer

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
		address: address,
		port:    port,
		weight:  weight,
		startAt: uint64(time.Now().UnixNano()),
	}
}
