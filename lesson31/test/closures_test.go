package test

import (
	"testing"
	"time"
)

func TestClosures(t *testing.T) {
	// x, y := 1, 2
	// func (a, b int) {
	// 	t.Log(a + b)
	// }(x, y)

	// for i := 0; i < 10; i++ {
	// 	sum := func () int {
	// 		return i + 1
	// 	}()
	// 	t.Logf("i = %d sum = %d \n", i, sum)
	// }

	// values := []int{1,2,3,4,5}
	// for _, val := range values {
	// 	go func() {
	// 		t.Logf("%p %d\n", &val, val)
	// 	}()
	// }

	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func(val int) {
			t.Logf("%p %d\n", &val, val)
		}(val)
	}
	time.Sleep(time.Second)
}
