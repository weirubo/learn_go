package lesson27

import (
	"testing"
)

func TestSum(t *testing.T) {
	a, b := 1, 2
	rst := Sum(a, b)
	if rst == 3 {
		t.Logf("expected=%d, actual=%d", 3, rst)
	} else {
		// t.Errorf("expected=%d, actual=%d", 3, rst)
		t.Fatalf("expected=%d, actual=%d", 3, rst)
	}
	t.Log("done")
}

func TestAvg(t *testing.T) {
	rst := Avg(1, 2, 3)
	if rst == 2 {
		t.Log("success")
		return
	}
	t.Log("fail")
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
