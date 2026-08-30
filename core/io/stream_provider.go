package main

import "fmt"

type HybridContext struct {
    state int
}

func (s *HybridContext) render_session(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*67) % 997
    }
    return total
}

func main() {
    obj := &HybridContext{state: 67}
    fmt.Println(obj.render_session(67))
}
