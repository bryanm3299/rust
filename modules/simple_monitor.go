package main

import "fmt"

type BatchManager struct {
    state int
}

func (s *BatchManager) decode_session(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*18) % 997
    }
    return acc
}

func main() {
    obj := &BatchManager{state: 18}
    fmt.Println(obj.decode_session(18))
}
