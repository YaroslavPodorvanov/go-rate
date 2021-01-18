package go_rate

import (
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
	"time"
)

type allow interface {
	Allow(id uint16, limit, now uint32) bool
}

const (
	maxIDs = 65535
	maxRPS = 64
)

func TestMapMutex_Allow(t *testing.T) {
	testAllow(t, NewMapMutex())
}

func TestArrayMutex_Allow(t *testing.T) {
	testAllow(t, NewArrayMutex())
}

func testAllow(t *testing.T, rate allow) {
	t.Helper()

	var now = now()

	require.Equal(t, true, rate.Allow(1, 1, now))
	require.Equal(t, true, rate.Allow(2, 1, now))
	require.Equal(t, true, rate.Allow(3, 1, now))
	require.Equal(t, false, rate.Allow(1, 1, now))
	require.Equal(t, false, rate.Allow(2, 1, now))
	require.Equal(t, false, rate.Allow(3, 1, now))

	require.Equal(t, true, rate.Allow(1, 2, now))
	require.Equal(t, true, rate.Allow(2, 2, now))
	require.Equal(t, true, rate.Allow(3, 2, now))
	require.Equal(t, false, rate.Allow(1, 2, now))
	require.Equal(t, false, rate.Allow(2, 2, now))
	require.Equal(t, false, rate.Allow(3, 2, now))

	require.Equal(t, true, rate.Allow(1, 2, now+1))
	require.Equal(t, true, rate.Allow(2, 2, now+1))
	require.Equal(t, true, rate.Allow(3, 2, now+1))
}

func BenchmarkMapMutex_Allow(b *testing.B) {
	benchmarkAllow(b, NewMapMutex())
}

func BenchmarkArrayMutex_Allow(b *testing.B) {
	benchmarkAllow(b, NewArrayMutex())
}

func benchmarkAllow(b *testing.B, rate allow) {
	b.Helper()

	var i uint32

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var index = uint16(atomic.AddUint32(&i, 1) & maxIDs)

			rate.Allow(index, maxRPS, now())
		}
	})
}

func now() uint32 {
	return uint32(time.Now().Unix())
}
