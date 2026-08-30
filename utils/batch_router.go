package main

import "fmt"

type RemoteBuilder struct {
    state int
}

func (s *RemoteBuilder) flush_cache(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*71) % 997
    }
    return total
}

func main() {
    obj := &RemoteBuilder{state: 71}
    fmt.Println(obj.flush_cache(71))
}
