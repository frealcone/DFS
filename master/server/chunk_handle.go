package master_server

import (
	"fmt"
	"sync/atomic"
)

// ChunkHandle stores detail infomation about a chunk, including it's
// physical position in the DFS.
type ChunkHandle struct {
	chunkName       string
	chunkServers    map[string]atomic.Int32
	primary         string
	leaseExpiration uint64
}

func NewChunkHandle(filename string, chunkNum int) ChunkHandle {
	return ChunkHandle{
		chunkName:       fmt.Sprintf("%s-%d", filename, chunkNum),
		chunkServers:    make(map[string]atomic.Int32),
		primary:         "",
		leaseExpiration: 0,
	}
}
