package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const defaultRootFolderName = "jynetwork"

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	Filename string
}

func (p *PathKey) FirstPathName() string {
	paths := strings.Split(p.PathName, string(filepath.Separator))
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

func (p *PathKey) FullPath() string {
	return filepath.Join(p.PathName, p.Filename)
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i+1)*blocksize
		paths[i] = hashStr[from:to]
	}
	return PathKey{
		PathName: filepath.Join(paths...),
		Filename: hashStr,
	}
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
	Root              string
}
type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if opts.Root == "" {
		opts.Root = defaultRootFolderName
	}
	return &Store{StoreOpts: opts}
}

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)
	_, err := os.Stat(filepath.Join(s.Root, pathKey.FullPath()))
	return !errors.Is(err, os.ErrNotExist)
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)
	defer func() {
		log.Printf("delete [%s] from disk", pathKey.Filename)
	}()
	return os.RemoveAll(filepath.Join(s.Root, pathKey.FirstPathName()))
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	return os.Open(filepath.Join(s.Root, pathKey.FullPath()))
}

func (s *Store) writerStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(filepath.Join(s.Root, pathKey.PathName), os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.FullPath()

	f, err := os.Create(filepath.Join(s.Root, pathKey.FullPath()))
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written %d bytes to disk: %s\n", n, fullPath)
	return nil
}
