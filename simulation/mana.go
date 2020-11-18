package simulation

import (
	"math"

	"github.com/iotaledger/hive.go/autopeering/mana"
	"github.com/iotaledger/hive.go/identity"
)

type manaFunc mana.Func

var manaF manaFunc

var IdentityMana map[*identity.Identity]uint64

func (f manaFunc) Eval(identity *identity.Identity) uint64 {
	return IdentityMana[identity]
}

func NewZipfMana(identities []*identity.Identity, zipf float64) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(identities))
	scalingFactor := math.Pow(10, 10)
	for i, p := range identities {
		m[p] = uint64(math.Pow(float64(i+1), -zipf) * scalingFactor)
		//log.Println(m[p])
	}
	return m
}

func NewZipfManaAdv(identities []*identity.Identity, zipf float64, nuNodes int, nuNodesAdv int, advManaPercent float64) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(identities))
	scalingFactor := math.Pow(10, 10)
	totalHonestMana := float64(0)
	for j := 0; j < nuNodes; j++ {
		totalHonestMana = totalHonestMana + math.Pow(float64(j+1), -zipf)*scalingFactor
	}
	totalAdvMana := advManaPercent / (1 - advManaPercent) * totalHonestMana
	for i, p := range identities {
		if i < nuNodes {
			m[p] = uint64(math.Pow(float64(i+1), -zipf) * scalingFactor)
		} else {
			m[p] = uint64(totalAdvMana / float64(nuNodesAdv))
		}
		//log.Println(m[p])
	}
	return m
}
