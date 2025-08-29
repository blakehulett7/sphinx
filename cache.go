package main

import (
	"fmt"
	"time"
)

var Cache = map[string]any{}

func CacheCleaner(interval time.Duration) {
	for {
		<-time.After(interval)

		fmt.Println()
		ColorPrint(Yellow, "---------------------------")
		ColorPrint(Yellow, fmt.Sprintf("Clearing cache... interval: %v", interval))
		ColorPrint(Yellow, "---------------------------")
		fmt.Println()

		Cache = map[string]any{}
	}
}

func WithCache[T any](key string, f func() T) T {
	cached_result, ok := Cache[key]
	if ok {
		return cached_result.(T)
	}

	result := f()
	Cache[key] = result

	return result
}
