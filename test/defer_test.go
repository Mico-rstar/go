package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeq(t *testing.T) {
	defer func ()  {
		fmt.Println("1")
	}()

	defer func ()  {
		fmt.Println("2")
	}()

	panic("")
}


func test1() int {
        i := 0
        defer func() {
                fmt.Println("defer1")
        }()
        defer func() {
                i += 1
                fmt.Println("defer2")
        }()
        return i
}

func test2() (i int) {
        i = 0
        defer func() {
                i += 1
                fmt.Println("defer2")
        }()
        return i
}

func TestExTime1(t *testing.T) {
	require.Equal(t, 0, test1())
}
func TestExTime2(t *testing.T) {
	require.Equal(t, 1, test2())
}