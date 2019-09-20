package neighborhood

import (
	"sync"

	"github.com/wollac/autopeering/distance"
	"github.com/wollac/autopeering/peer"
)

type Neighborhood struct {
	Neighbors []peer.PeerDistance
	Size      int
	mutex     sync.RWMutex
}

func (nh *Neighborhood) getFurtherest() (peer.PeerDistance, int) {
	nh.mutex.RLock()
	defer nh.mutex.RUnlock()
	if len(nh.Neighbors) < nh.Size {
		return peer.PeerDistance{
			Remote:   nil,
			Distance: distance.Max,
		}, len(nh.Neighbors)
	}

	index := 0
	furtherest := nh.Neighbors[index]
	for i, neighbor := range nh.Neighbors {
		if neighbor.Distance > furtherest.Distance {
			furtherest = neighbor
			index = i
		}
	}
	return furtherest, index
}

func (nh *Neighborhood) Select(candidates []peer.PeerDistance) peer.PeerDistance {
	if len(candidates) > 0 {
		target, _ := nh.getFurtherest()
		for _, candidate := range candidates {
			if candidate.Distance < target.Distance {
				return candidate
			}
		}
	}
	return peer.PeerDistance{}
}

func (nh *Neighborhood) Add(toAdd peer.PeerDistance) {
	nh.mutex.Lock()
	defer nh.mutex.Unlock()
	if len(nh.Neighbors) < nh.Size {
		nh.Neighbors = append(nh.Neighbors, toAdd)
	}
}

func (nh *Neighborhood) RemovePeer(toRemove peer.ID) {
	index := nh.getPeerIndex(toRemove)
	if index < 0 {
		return
	}
	nh.mutex.Lock()
	defer nh.mutex.Unlock()
	nh.Neighbors[index] = peer.PeerDistance{}
	copy(nh.Neighbors[index:], nh.Neighbors[index+1:])
	nh.Neighbors = nh.Neighbors[:len(nh.Neighbors)-1]
}

func (nh *Neighborhood) getPeerIndex(target peer.ID) int {
	nh.mutex.RLock()
	defer nh.mutex.RUnlock()
	for i, peer := range nh.Neighbors {
		if peer.Remote.ID() == target {
			return i
		}
	}
	return -1

}

func (nh *Neighborhood) UpdateDistance(anchor, salt []byte) {
	nh.mutex.Lock()
	defer nh.mutex.Unlock()
	for i, peer := range nh.Neighbors {
		nh.Neighbors[i].Distance = distance.BySalt(anchor, peer.Remote.ID().Bytes(), salt)
	}
}

func (nh *Neighborhood) GetPeers() []*peer.Peer {
	nh.mutex.RLock()
	defer nh.mutex.RUnlock()
	list := make([]*peer.Peer, len(nh.Neighbors))
	for i, peer := range nh.Neighbors {
		list[i] = peer.Remote
	}
	return list
}

func (nh *Neighborhood) GetPeerFromID(id peer.ID) *peer.Peer {
	nh.mutex.RLock()
	defer nh.mutex.RUnlock()
	for _, peer := range nh.Neighbors {
		if peer.Remote.ID() == id {
			return peer.Remote
		}
	}
	return nil
}