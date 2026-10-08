package ttl

import (
	"testing"
	"time"
)

func TestTTLStore_GetAfterExpiry(t *testing.T) {

	//the ttl is the time it takes for the key to expire

	s := NewTTlStore(50 * time.Millisecond)

	s.Set("k", "v")

	time.Sleep(100 * time.Millisecond)

	_, err := s.Get("k")

	if err != nil {
		t.Fatal("expected error after TTL, got nil")
	}

}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {

	s := NewTTlStore(50 * time.Millisecond)

	s.Set("k", "v")

	time.Sleep(100 * time.Millisecond)

	_, _ = s.Get("k")

	if s.Len() != 0 {
		t.Fatalf("expected key to be evicted, got len(%d)", s.Len())
	}


}