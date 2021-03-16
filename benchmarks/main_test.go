package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}
func BenchmarkSet10_90(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixNano())
	var ds = make([]int, 1e8)
	for i := 0; i < 1e8; i++ {
		ds[i] = rand.Intn(100)
	}

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {

			i := 0
			for pb.Next() {
				if ds[i] < 10 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}
func BenchmarkSet50_50(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixNano())
	var ds = make([]int, 1e8)
	for i := 0; i < 1e8; i++ {
		ds[i] = rand.Intn(100)
	}
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if ds[i] < 50 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

func BenchmarkSet90_10(b *testing.B) {
	var set = NewSet()
	rand.Seed(time.Now().UnixNano())
	var ds = make([]int, 1e8)
	for i := 0; i < 1e8; i++ {
		ds[i] = rand.Intn(100)
	}
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if ds[i] < 90 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}
