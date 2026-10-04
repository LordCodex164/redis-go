package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore(2)
	store.Set("charlie", "3")
	store.Set("james", "1")
	store.Set("bravo", "2")

	got := store.Keys()
	want := []string{"alpha", "bravo", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {

	store := NewStore(2)

	got := store.Keys()

	//empty keys

	fmt.Println("got", got)

	//want an
	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(len(want), len(got)) {
		t.Errorf("The Keys is not empty")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	//CHALLENGE
	//1. Make a NewStore.
	//2. Set("hello", "world")

	store := NewStore(2)

	store.Set("hello", "world")

	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() failed")
	}

	if val, err := store.Get("missing"); err != nil || val != "world" {
		t.Errorf("Get() failed")
	}

}

func TestDeleteExistingKey(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"create", map[string]string{"a": "1", "b": "4"}, "a", 1}, //pass
		// {"delete", map[string]string{"b": "2", "c": "3"}, "c", 2}, //fail
		// {"update", map[string]string{"b": "2", "p": "4"}, "v", 1}, //fail
		// {"read", map[string]string{"b": "2"}, "b", 3},             //fail
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore(2)
			for k, v := range tc.setup {
				store.Set(k, v)
				fmt.Println("key", tc.deleteKey)
			}
			store.Delete(tc.deleteKey)
			fmt.Println("update_keys", store.Keys())
			if got := store.Keys(); tc.wantLen != store.Len() {
				t.Errorf("delete(%q) Length is (%v), want (%d)", tc.deleteKey, len(got), tc.wantLen)
			}
		})
	}
}

type test struct {

}

func TestRenameKey(t *testing.T) {
	store := NewStore(2)

	tests := []struct{
		name 		string
		obj			map[string]string
		deleteKey	string
		wantLength	int
	}{
		{"create", map[string]string{"a": "1", "b": "2"}, "a", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			store := NewStore(2)

			for k, v := range test.obj {
				store.Set(k, v)
			}

			store.Rename("a", "c")

			fmt.Println("len", store.Len())

			if l := store.Len(); l != test.wantLength {
				t.Errorf("Expected Length: %d got %v", test.wantLength, store.Len())
			}
		})
	}

	store.Set("dns", "lookup")
	store.Set("dn", "phonebook")
}