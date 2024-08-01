package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	// x.TestRedis()
	fmt.Println(prometheus.ExponentialBuckets(0.00025, 10, 5))
}
