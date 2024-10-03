package master_server_test

import (
	"container/heap"
	"testing"

	server "github.com/frealcone/DFS/master/server"
)

func TestLocalRegistry(t *testing.T) {
	cases := make([][2][]server.ChunkServer, 1)
	cases[0] = [2][]server.ChunkServer{
		{
			server.NewChunkServer("192.168.19.159", 2222, 3),
			server.NewChunkServer("192.168.19.169", 3333, 1),
			server.NewChunkServer("192.168.220.160", 4000, 1),
			server.NewChunkServer("178.20.199.175", 3000, 2),
		},
		{
			server.NewChunkServer("192.168.19.159", 2222, 3),
			server.NewChunkServer("178.20.199.175", 3000, 2),
			server.NewChunkServer("192.168.19.169", 3333, 1),
			server.NewChunkServer("192.168.220.160", 4000, 1),
		},
	}

	for _, qa := range cases {
		qs := server.LocalRegistry(qa[0])

		heap.Init(&qs)

		for i, a := range qa[1] {
			if q := heap.Pop(&qs).(server.ChunkServer); q.Compare(a) != 0 {
				t.Fatalf("%dth test case failed:\nYour output:\n%v\nCorrect output:\n%v\n", i, q, a)
			}
		}
	}
}
