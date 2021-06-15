package code_engine

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"sort"
	"testing"
)

func TestCodeEngine_MergeMultiArrays(t *testing.T) {
	convey.Convey("test code engine merge multi sorted arrays", t, func() {
		// merge arrays test
		input := [][]int{{1, 3, 5}, {2, 4, 6}}
		checkRes := mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{}, {}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{}, {}, {}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1}, {}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1}, {1, 2}, {3, 9, 12}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 1, 1}, {2, 2, 2, 2, 2}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1}, {}, {-1, 7, 9}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {2, 4, 6}, {-1, 7, 9, 15, 21}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {2, 4, 6}, {-1, 7, 9}, {-5}, {10, 20, 30, 50, 70}, {100}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {2, 4, 6}, {-1, 7, 9}, {-5}, {10, 20, 30, 50, 70}, {100}}
		input = append(input, input...)
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {2, 4, 6}, {-1, 7, 9}, {-5}, {10, 20, 30, 50, 70}, {100}, {}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1, 3, 5}, {2, 4, 6}, {-1, 7, 9}, {-5}, {10, 20, 30, 50, 70}, {100}, {-5, -1, 1, 2, 10, 100}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, true)

		input = [][]int{{1}, {}, {-1, 7, 5}}
		checkRes = mergeTestCase(input)
		convey.So(checkRes, convey.ShouldEqual, false)
	})
}

func mergeTestCase(input [][]int) bool {
	fmt.Printf("============================\n")
	fmt.Printf("inputs: %v\n", input)
	expected := make([]int, 0, len(input))
	for _, data := range input {
		t := make([]int, len(data), len(data))
		copy(t, data)
		expected = append(expected, t...)
	}
	sort.Ints(expected)

	ans := MergeMultiSortedArrays(input)
	fmt.Printf("ans : %v\n", ans)

	return checkResEqual(ans, expected)
}

func BenchmarkCodeEngine_MergeMultiArraysWithSimpleHeap(b *testing.B) {
	convey.Convey("test benchmark code engine merge multi sorted arrays with simple heap", b, func() {
		// 256-pathway, 256w total, about 0.75 seconds except init time
		multi := 256
		size := 10000
		max := 100000000

		var input [][]int
		for i := 0; i < multi; i++ {
			ia := make([]int, 0, size)
			for i := 0; i < size; i++ {
				ia = append(ia, rand.Intn(max))
			}
			sort.Ints(ia)
			input = append(input, ia)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ans := MergeMultiSortedArrays(input)
			convey.So(len(ans), convey.ShouldEqual, size*multi)
			//convey.So(sort.IntsAreSorted(ans), convey.ShouldEqual, true)
		}
	})
}

func BenchmarkCodeEngine_MergeMultiArraysWithStdHeap(b *testing.B) {
	convey.Convey("test benchmark code engine merge multi sorted arrays with std heap", b, func() {
		// 256-pathway, 256w total, about 0.5 seconds except init time
		multi := 256
		size := 10000
		max := 100000000

		var input [][]int
		for i := 0; i < multi; i++ {
			ia := make([]int, 0, size)
			for i := 0; i < size; i++ {
				ia = append(ia, rand.Intn(max))
			}
			sort.Ints(ia)
			input = append(input, ia)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ans := MergeMultiSortedArraysWithStdHeap(input)
			convey.So(len(ans), convey.ShouldEqual, size*multi)
			//convey.So(sort.IntsAreSorted(ans), convey.ShouldEqual, true)
		}
	})
}

func BenchmarkCodeEngine_MergeMultiArraysWithSimpleHeapFewPaths(b *testing.B) {
	convey.Convey("test benchmark code engine merge multi sorted arrays with simple heap", b, func() {
		// 32-pathway, 320w total, about 0.26 seconds except init time
		multi := 32
		size := 100000
		max := 100000000

		var input [][]int
		for i := 0; i < multi; i++ {
			ia := make([]int, 0, size)
			for i := 0; i < size; i++ {
				ia = append(ia, rand.Intn(max))
			}
			sort.Ints(ia)
			input = append(input, ia)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ans := MergeMultiSortedArrays(input)
			convey.So(len(ans), convey.ShouldEqual, size*multi)
			//convey.So(sort.IntsAreSorted(ans), convey.ShouldEqual, true)
		}
	})
}

func BenchmarkCodeEngine_MergeMultiArraysWithStdHeapFewPaths(b *testing.B) {
	convey.Convey("test benchmark code engine merge multi sorted arrays with std heap", b, func() {
		// 32-pathway, 320w total, about 0.4 seconds except init time
		multi := 32
		size := 100000
		max := 100000000

		var input [][]int
		for i := 0; i < multi; i++ {
			ia := make([]int, 0, size)
			for i := 0; i < size; i++ {
				ia = append(ia, rand.Intn(max))
			}
			sort.Ints(ia)
			input = append(input, ia)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ans := MergeMultiSortedArraysWithStdHeap(input)
			convey.So(len(ans), convey.ShouldEqual, size*multi)
			//convey.So(sort.IntsAreSorted(ans), convey.ShouldEqual, true)
		}
	})
}
