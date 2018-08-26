package uintset

import (
	"testing"

	"github.com/shenwei356/intintmap"
)

func TestMapSimple(t *testing.T) {
	m := New(10, 0.99)
	var i uint64

	// --------------------------------------------------------------------
	// Add() and Has()

	for i = 0; i < 20000; i += 2 {
		m.Add(i)
	}
	for i = 0; i < 20000; i += 2 {
		if !m.Has(i) {
			t.Errorf("%d should exist", i)
		}
		if m.Has(i + 1) {
			t.Errorf("%d should not exist", i)
		}
	}

	if m.Size() != int(20000/2) {
		t.Errorf("size (%d) is not right, should be %d", m.Size(), int(20000/2))
	}

	// --------------------------------------------------------------------
	// Keys()

	m0 := make(map[uint64]uint64, 1000)
	for i = 0; i < 20000; i += 2 {
		m0[i] = i
	}
	n := len(m0)

	for k := range m.Items() {
		m0[k] = k * 2
	}
	if n != len(m0) {
		t.Errorf("get unexpected more keys")
	}

	for k, v := range m0 {
		if k*2 != v {
			t.Errorf("didn't get expected changed value")
		}
	}

	// --------------------------------------------------------------------
	// Del()

	for i = 0; i < 20000; i += 2 {
		m.Del(i)
	}
	for i = 0; i < 20000; i += 2 {
		if m.Has(i) {
			t.Errorf("%d should not exist", i)
		}
		if m.Has(i + 1) {
			t.Errorf("%d should not exist", i)
		}
	}

}

const MAX uint64 = 999999999
const STEP uint64 = 9534

func fillUintset(m *Set) {
	var j, k uint64
	for j = 0; j < MAX; j += STEP {
		m.Add(j)
		for k = j; k < j+16; k++ {
			m.Add(k)
		}
	}
}

func fillIntIntMap(m *intintmap.Map) {
	var j, k int64
	for j = 0; j < int64(MAX); j += int64(STEP) {
		m.Put(j, 0)
		for k = j; k < j+16; k++ {
			m.Put(k, 0)
		}
	}
}

func fillStdMap(m map[uint64]struct{}) {
	var j, k uint64
	for j = 0; j < MAX; j += STEP {
		m[j] = struct{}{}
		for k = j; k < j+16; k++ {
			m[k] = struct{}{}
		}
	}
}

func BenchmarkUintsetFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := New(2048, 0.80)
		fillUintset(m)
	}
}

func BenchmarkIntIntMapFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := intintmap.New(2048, 0.80)
		fillIntIntMap(m)
	}
}

func BenchmarkStdMapFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[uint64]struct{}, 2048)
		fillStdMap(m)
	}
}

func BenchmarkUintsetTest10PercentHitRate(b *testing.B) {
	var j, k, sum uint64
	m := New(2048, 0.80)
	fillUintset(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < MAX; j += STEP {
			for k = j; k < 10; k++ {
				if m.Has(k) {
					sum += k
				}
			}
		}
		//log.Println("int int sum:", sum)
	}
}

func BenchmarkIntIntMapTest10PercentHitRate(b *testing.B) {
	var j, k, sum int64
	var ok bool
	m := intintmap.New(2048, 0.80)
	fillIntIntMap(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < int64(MAX); j += int64(STEP) {
			for k = j; k < 10; k++ {
				if _, ok = m.Get(k); ok {
					sum += k
				}
			}
		}
		//log.Println("int int sum:", sum)
	}
}

func BenchmarkStdMapTest10PercentHitRate(b *testing.B) {
	var j, k, sum uint64
	var ok bool
	m := make(map[uint64]struct{}, 2048)
	fillStdMap(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < MAX; j += STEP {
			for k = j; k < 10; k++ {
				if _, ok = m[k]; ok {
					sum += k
				}
			}
		}
		//log.Println("map sum:", sum)
	}
}

func BenchmarkUintsetTest100PercentHitRate(b *testing.B) {
	var j, sum uint64
	m := New(2048, 0.80)
	fillUintset(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < MAX; j += STEP {
			if m.Has(j) {
				sum += j
			}
		}
		//log.Println("int int sum:", sum)
	}
}

func BenchmarkIntIntMapTest100PercentHitRate(b *testing.B) {
	var j, sum int64
	var ok bool
	m := intintmap.New(2048, 0.80)
	fillIntIntMap(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < int64(MAX); j += int64(STEP) {
			if _, ok = m.Get(j); ok {
				sum += j
			}
		}
		//log.Println("int int sum:", sum)
	}
}

func BenchmarkStdMapTest100PercentHitRate(b *testing.B) {
	var j, sum uint64
	var ok bool
	m := make(map[uint64]struct{}, 2048)
	fillStdMap(m)
	for i := 0; i < b.N; i++ {
		sum = 0
		for j = 0; j < MAX; j += STEP {
			if _, ok = m[j]; ok {
				sum += j
			}
		}
		//log.Println("map sum:", sum)
	}
}
