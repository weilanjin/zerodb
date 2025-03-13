package main

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "weijunyu"
	pathKey := CASPathTransformFunc(key)
	originalKey := "4652a438f48a0b6b52ddfa9a713a673f987dcb57"
	expectedPathName := "4652a/438f4/8a0b6/b52dd/fa9a7/13a67/3f987/dcb57"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s, want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != originalKey {
		t.Errorf("have %s, want %s", pathKey.Filename, originalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	store := NewStore(opts)

	key := "weijunyu"
	data := "hello world"
	if err := store.writerStream(key, bytes.NewBufferString(data)); err != nil {
		t.Error(err)
	}
	if err := store.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	store := NewStore(opts)

	key := "weijunyu"
	data := "hello world"
	if err := store.writerStream(key, bytes.NewBufferString(data)); err != nil {
		t.Error(err)
	}

	if !store.Has(key) {
		t.Errorf("expected key %s to exist", key)
	}

	r, err := store.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	log.Println(string(b))

	if string(b) != data {
		t.Errorf("have %s, want %s", string(b), data)
	}
	store.Delete(key)
}
