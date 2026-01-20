package main

import (
	"errors"
	"sync"
	"time"
)

// 一道考题
// 要求实现一个 map:
// (1)面向高并发;
// (2)只存在插入和查询操作 0(1);
// (3)查询时，若 key 存在，直接返回 val;若 key 不存在，阻塞直到 key·val 对被放入后，获取 val返回;等待指定时长仍未放入，返回超时错误;
// (4)写出真实代码，不能有死锁或者 panic 风险.

type MyConcurrentMap struct {
	inner map[int]int
	sync.Mutex
	signal map[int]chan struct{}
}

func New() *MyConcurrentMap {
	return &MyConcurrentMap{
		inner: make(map[int]int),
		signal: make(map[int]chan struct{}),
	}
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	// map中已经存在k
	if v, ok := m.inner[k]; ok {
		return v, nil
	}
	// map不存在k，等待chan信号，直到超时
	if _, ok := m.signal[k]; !ok {
		// 只有第一次创建channel
		m.signal[k] = make(chan struct{})
	}
	m.Unlock()
	select {
	case <-m.signal[k]:
		m.Lock()
		res := m.inner[k]
		m.Unlock()
		return res, nil
	case <-time.After(maxWaitingDuration):
		return 0, errors.New("overtime")
	}
}

func (m *MyConcurrentMap) Put(k int, v int) {
	m.inner[k] = v
	if c, ok := m.signal[k]; ok {
		// 存在chan，说明有reader需要唤醒
		// c <- struct{}{}
		// 这里用close是因为要唤醒多个reader，使用close模拟广播机制
		close(c)
	}
}

func main() {

}
