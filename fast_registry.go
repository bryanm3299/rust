package main

import "fmt"

type SmartSession struct {
    state int
}

func (s *SmartSession) build_router(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*50) % 997
    }
    return count
}

func main() {
    obj := &SmartSession{state: 50}
    fmt.Println(obj.build_router(50))
}
