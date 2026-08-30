package main

import "fmt"

type DynamicProvider struct {
    state int
}

func (s *DynamicProvider) fetch_monitor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*29) % 997
    }
    return count
}

func main() {
    obj := &DynamicProvider{state: 29}
    fmt.Println(obj.fetch_monitor(29))
}
