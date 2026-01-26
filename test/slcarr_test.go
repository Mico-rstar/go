package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArr2Slice(t *testing.T) {
	// 等价写法
	// a := [...]int{1, 2, 3}
	// a := [3]int{1, 2, 3}
	a := [...]int{1, 2, 3}
	// a[1, 2)
	slc := a[1:2]
	require.Equal(t, len(slc), 1)
	// cap = len(a) - 1
	require.Equal(t, cap(slc), 2)

}

func TestSlcAppend1(t *testing.T) {
	a := [...]int{1, 2, 3}
	slc := a[0:1]
	slc[0] = 3
	require.Equal(t, 3, a[0])
	slc = append(slc, 1)
	require.Equal(t, 1, slc[1])
	require.Equal(t, a[1], 1)
	// a = {3, 1, 3}
	// slc = {3, 1}

	// 扩容到超出容量，底层数组发生改变
	slc = append(slc, 2, 3)
	require.Equal(t, 2, slc[2])
	require.Equal(t, 3, slc[3])
	require.Equal(t, a[2], 3)

}

func TestSlcAppend2(t *testing.T) {
	a := [...]int{1, 2, 3}
	// [low, high, max]
	// cap = max - low
	s1 := a[0:1:2]
	require.Equal(t, 1, len(s1))
	require.Equal(t, 2, cap(s1))

	// panic: runtime error: slice bounds out of range
	// s2 := a[1:4]
	// require.Equal(t, 3, len(s2))
}

func TestSlcAppend3(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5] // {2, 3, 4}
	s2 := s1[2:6:7]  // {4, 5, 6, 7}
	require.Equal(t, 4, len((s2)))
	require.Equal(t, 3, len(s1))

	// len = high - low
	// cap = max - low
	// panic
	high := 7
	s3 := s1[2:high:6]
	require.Equal(t, 4, len(s3))

}
