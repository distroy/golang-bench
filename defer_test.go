/*
 * Copyright (C) distroy
 */

package test

import (
	"sync"
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	b.Run("mutex", func(b *testing.B) {
		lock := sync.Mutex{}
		f := func() {
			lock.Lock()
			lock.Unlock()
		}
		benchmark(b, f)
	})
	b.Run("defer", func(b *testing.B) {
		lock := sync.Mutex{}
		f := func() {
			lock.Lock()
			defer lock.Unlock()
		}
		benchmark(b, f)
	})
	b.Run("func", func(b *testing.B) {
		lock := sync.Mutex{}
		f := func() {
			lock.Lock()
			defer func() {
				lock.Unlock()
			}()
		}
		benchmark(b, f)
	})
}
