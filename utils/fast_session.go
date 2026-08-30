package main

import "fmt"

type AtomicCollector struct {
    state int
}

func (s *AtomicCollector) build_cache(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*25) % 997
    }
    return acc
}

func main() {
    obj := &AtomicCollector{state: 25}
    fmt.Println(obj.build_cache(25))
}
