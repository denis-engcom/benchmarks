package groups_and_pools_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"golang.org/x/sync/errgroup"
)

func BenchmarkStdWaitGroup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 1000
		count := atomic.Int32{}
		wg := sync.WaitGroup{}
		wg.Add(sum)
		for j := 0; j < sum; j++ {
			go func() {
				defer wg.Done()
				count.Add(1)
			}()
		}
		wg.Wait()
		if count.Load() != int32(sum) {
			b.Fatal("unexpected sum")
		}
	}
}

func BenchmarkStdErrgroup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 1000
		count := atomic.Int32{}
		g := errgroup.Group{}
		for j := 0; j < sum; j++ {
			g.Go(func() error {
				count.Add(1)
				return nil
			})
		}
		_ = g.Wait()
		if count.Load() != int32(sum) {
			b.Fatal("unexpected sum")
		}
	}
}

func BenchmarkConcWaitGroup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 1000
		count := atomic.Int32{}
		wg := conc.WaitGroup{}
		for j := 0; j < sum; j++ {
			wg.Go(func() {
				count.Add(1)
			})
		}
		wg.Wait()
		if count.Load() != int32(sum) {
			b.Fatal("unexpected sum")
		}
	}
}

func BenchmarkConcErrorPool1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 1000
		count := atomic.Int32{}
		pool := pool.New().WithErrors()
		for j := 0; j < sum; j++ {
			pool.Go(func() error {
				count.Add(1)
				return nil
			})
		}
		_ = pool.Wait()
		if count.Load() != int32(sum) {
			b.Fatal("unexpected sum")
		}
	}
}

func BenchmarkConcContextPool1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 1000
		count := atomic.Int32{}
		pool := pool.New().
			WithContext(context.Background()).
			WithCancelOnError().
			WithFirstError()
		for j := 0; j < sum; j++ {
			pool.Go(func(_ context.Context) error {
				count.Add(1)
				return nil
			})
		}
		_ = pool.Wait()
		if count.Load() != int32(sum) {
			b.Fatal("unexpected sum")
		}
	}
}
