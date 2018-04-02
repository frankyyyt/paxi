package wpaxos

import (
	"github.com/ailidani/paxi"
	"github.com/ailidani/paxi/paxos"
)

type kpaxos struct {
	paxi.Node
	key paxi.Key
	*paxos.Paxos
	paxi.Policy
}

func newKPaxos(key paxi.Key, node paxi.Node) *kpaxos {
	k := &kpaxos{}
	k.Node = node
	k.key = key
	k.Policy = paxi.NewPolicy()
	k.Paxos = paxos.NewPaxos(k)
	return k
}

// Broadcast overrides Socket interface in Node
func (k *kpaxos) Broadcast(m interface{}) {
	switch m := m.(type) {
	case paxos.P1a:
		k.Node.Broadcast(Prepare{k.key, m})
	case paxos.P2a:
		k.Node.Broadcast(Accept{k.key, m})
	case paxos.P3:
		k.Node.Broadcast(Commit{k.key, m})
	default:
		k.Node.Broadcast(m)
	}
}

// Send overrides Socket interface in Node
func (k *kpaxos) Send(to paxi.ID, m interface{}) {
	switch m := m.(type) {
	case paxos.P1b:
		k.Node.Send(to, Promise{k.key, m})
	case paxos.P2b:
		k.Node.Send(to, Accepted{k.key, m})
	default:
		k.Node.Send(to, m)
	}
}