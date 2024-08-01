package x

import (
	"fmt"
	"math"
	"sync/atomic"
	"testing"
	"time"

	"github.com/coredns/coredns/plugin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func TestNewHistogramObserve(t *testing.T) {
	var bits uint64
	v := 2.4
	bits = math.Float64bits(12)
	loadedBits := atomic.LoadUint64(&bits)
	v2 := math.Float64frombits(loadedBits) + v
	newBits := math.Float64bits(v2)
	fmt.Println(&newBits)
	flag := atomic.CompareAndSwapUint64(&bits, loadedBits, newBits)
	fmt.Println(flag)

	loadedBits = atomic.LoadUint64(&bits)
	v2 = math.Float64frombits(loadedBits) + v
	newBits = math.Float64bits(v2)
	fmt.Println(&newBits)
	flag = atomic.CompareAndSwapUint64(&bits, loadedBits, newBits)
	fmt.Println(flag)

	start := time.Now()
	HealthDuration := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace:                   "coredns",
		Subsystem:                   "health",
		Name:                        "request_duration_seconds",
		Buckets:                     plugin.SlimTimeBuckets,
		NativeHistogramBucketFactor: plugin.NativeHistogramBucketFactor,
		Help:                        "Histogram of the time (in seconds) each request took.",
	})
	HealthDuration.Observe(time.Since(start).Seconds())
	HealthDuration.Observe(time.Since(start).Seconds())
}
