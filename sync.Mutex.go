package main

import (
	"fmt"
	"sync"
	"time"
)

// type Test struct {
// 	mutex sync.Mutex
// 	v     int
// }

// func increment(test *Test) {
// 	test.mutex.Lock()
// 	defer test.mutex.Unlock()
// 	for i := 0; i < 10; i++ {
// 		test.v += 1
// 		fmt.Println("one:", test.v)
// 	}
// }

// func decrement(test *Test) {
// 	test.mutex.Lock()
// 	defer test.mutex.Unlock()
// 	fmt.Println("Is it locked?")
// 	for i := 0; i < 10; i++ {
// 		test.v -= 1
// 		fmt.Println(test.v)
// 	}
// 	fmt.Println("Is it unlocked?")
// }

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key]++
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func main() {
	// test := Test{v: 0}
	// go increment(&test)
	// go decrement(&test)
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("some")
	}
	time.Sleep(time.Second)

	fmt.Println(c.Value("some"))
}

// https://go.dev/tour/concurrency/9
