package go_rate

import (
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxIDs = 65535
	maxRPS = 64
)

func TestMapMutex_Allow(t *testing.T) {
	var rate = NewMapMutex()

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
	var rate = NewMapMutex()

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
