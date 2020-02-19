package basis

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Reduce(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.v[key]--
}

func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

func Run_mutex() {
	s := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.Inc("somekey")
		go s.Reduce("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Value("somekey"))
}
