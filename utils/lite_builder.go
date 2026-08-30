package main

import "fmt"

type AsyncCache struct {
    state int
}

func (s *AsyncCache) resolve_cache(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*90) % 997
    }
    return total
}

func main() {
    obj := &AsyncCache{state: 90}
    fmt.Println(obj.resolve_cache(90))
}
