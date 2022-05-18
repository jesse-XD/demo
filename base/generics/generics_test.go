package generics

import "testing"

func TestGMin(t *testing.T) {
	var (
		a int64 = 1
		b int64 = 10
	)
	m := GMin[int64](a, b)
	if m != a {
		t.Fatal("GMin int64 fatal...")
	}
	var (
		a2 float64 = 1.11
		b2 float64 = 22.3
	)
	m2 := GMin[float64](a2, b2)
	if m2 != a2 {
		t.Fatal("GMin float64 fatal...")
	}
}

func TestGSum(t *testing.T) {
	ints := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	floats := map[string]float64{
		"a": 1.2,
		"b": 2.4,
		"c": 3.6,
	}
	t.Logf("ints sum %v; floats sum %v", GSum[string, int64](ints), GSum[string, float64](floats))
}

func TestGSum2(t *testing.T) {
	ints := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	floats := map[string]float64{
		"a": 1.2,
		"b": 2.4,
		"c": 3.6,
	}
	// 另一种写法
	iGSum := GSum[string, int64]
	fGSum := GSum[string, float64]
	t.Logf("ints sum %v; floats sum %v", iGSum(ints), fGSum(floats))
}
