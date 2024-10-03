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

func TestRegister(t *testing.T) {
	registry := new(server.LocalRegistry)

	registry.Register(server.NewChunkServer("192.168.17.199", 3301, 3))
	registry.Register(server.NewChunkServer("192.168.19.211", 3302, 5))
	registry.Register(server.NewChunkServer("192.168.20.220", 3303, 4))
	registry.Register(server.NewChunkServer("192.168.17.199", 3301, 4))

	if registry.Len() != 3 {
		t.Fatalf("duplicated chunk server infomation in registry: %v", registry)
	}
}

func TestNext(t *testing.T) {
	registry := new(server.LocalRegistry)

	w5server := server.NewChunkServer("192.168.19.211", 3302, 5)
	w3server := server.NewChunkServer("192.168.17.199", 3301, 3)
	w4server := server.NewChunkServer("192.168.20.220", 3303, 4)

	registry.Register(w3server)
	registry.Register(w5server)
	registry.Register(w4server)

	answers := []server.ChunkServer{
		w5server, w5server, w5server, w5server, w5server,
		w4server, w4server, w4server, w4server,
		w3server, w3server, w3server,
		w5server,
	}

	for i, ans := range answers {
		if next := registry.Next(); next.GetIPAddress() != ans.GetIPAddress() {
			t.Fatalf("%dth test case failed:\nYour output:\n%v\nCorrect output:\n%v\n", i, *next, ans)
		}
	}
}

func TestAll(t *testing.T) {
	registry := new(server.LocalRegistry)

	registry.Register(server.NewChunkServer("192.168.17.199", 3301, 3))
	registry.Register(server.NewChunkServer("192.168.19.211", 3302, 5))
	registry.Register(server.NewChunkServer("192.168.20.220", 3303, 4))
	registry.Register(server.NewChunkServer("192.168.17.199", 3301, 4))

	answers := map[string]struct{}{
		"192.168.17.199:3301": {},
		"192.168.19.211:3302": {},
		"192.168.20.220:3303": {},
	}

	for registry.Len() > 0 {
		ans := heap.Pop(registry).(server.ChunkServer).GetIPAddress()
		if _, ok := answers[ans]; !ok {
			t.Fatalf("invalid ip address: %s", ans)
			break
		}
		delete(answers, ans)
	}

	if len(answers) > 0 {
		t.Fatalf("ignored chunk servers: %v", answers)
	}
}
