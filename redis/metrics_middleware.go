package main

import (
	"fmt"
	"redis-app/store"
	"time"
)

type MetricsMiddleware struct {
	innerStore 		store.IStore
	getCalls		int
	setCalls		int
	deleteCalls		int
	missedCalls		int
	lenCalls		int
	totalSetLatency	time.Duration
	totalGetLatency	time.Duration
}

func NewMetricsMiddleware(store store.IStore) *MetricsMiddleware {
	return &MetricsMiddleware{
		innerStore: store,
		getCalls: 			0,
		setCalls: 			0,
		deleteCalls: 		0,
		missedCalls: 		0,
		totalSetLatency:	0,
		totalGetLatency: 	0,
	}
}

func (m *MetricsMiddleware) Set(key, val string) error {
	start := time.Now()
	if err := m.innerStore.Set(key, val); err != nil {
		return err
	}
	m.setCalls++
	m.totalSetLatency += time.Since(start)
	return nil
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()
	m.getCalls++
	val, err := m.innerStore.Get(key);
	if err != nil {
		m.missedCalls--
		return "", fmt.Errorf("")
	}
	m.totalGetLatency += time.Since(start)
	return val, nil
}

func (m *MetricsMiddleware) Delete(key string) (error) {
	m.deleteCalls++
	_, ok := m.innerStore.Pop(key)
	if !ok {
		return fmt.Errorf("missing key")
	}
	return nil
}

func (m *MetricsMiddleware) Len() int {
	m.lenCalls++
	return m.innerStore.Len()
}

func (m *MetricsMiddleware) Report() {
	fmt.Printf("GET calls: %d (Missed Calls: %d)\n", m.getCalls, m.missedCalls)
	fmt.Printf("DELETE calls: %d\n", m.totalGetLatency)

	if m.getCalls > 0 || m.totalGetLatency > 0 {
		fmt.Printf("Average calls: %d\n", m.totalGetLatency / time.Duration(m.getCalls))
	}
}