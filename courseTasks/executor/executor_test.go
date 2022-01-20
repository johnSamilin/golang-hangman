package executor

import (
	"strings"
	"sync"
	"testing"
)

func Benchmark_startExecution(b *testing.B) {
	var wg sync.WaitGroup
	cExit := make(chan bool)
	for i := 0; i < b.N; i++ {
		commands := strings.Split(strings.Repeat("TIME TIME TIME_WAIT TIME", i+1), " ")
		workers := make(chan Worker, b.N)
		for j := 0; j < b.N; j++ {
			workers <- Worker{Id: j, Lock: &sync.Mutex{}}
		}
		wg.Add(1)
		startExecution(b.N, workers, commands, cExit, &wg)
		close(workers)
	}
	wg.Wait()
}