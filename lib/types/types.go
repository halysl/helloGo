package types

import "sync"

type ElectionProof struct {
	WinCount int64
	VRFProof []byte
}

type ProofWithLock struct {
	Value int64
	Lock sync.RWMutex
}
