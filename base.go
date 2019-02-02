/*
 * Copyright (C) distroy
 */

package test

import (
	"testing"
)

type benchmarkHandler func()

func benchmark(b *testing.B, f benchmarkHandler) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f()
		}
	})
	b.ReportAllocs()
}
