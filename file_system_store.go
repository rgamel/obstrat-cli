package obstrat

import (
	"fmt"
	"os"
)

type FileSystemStratStore struct {
	lines []string
}

func NewFileSystemStratStore(file *os.File) (*FileSystemStratStore, error) {
	lines, err := GetStratsFromFile(file)

	if err != nil {
		return nil, err
	}

	return &FileSystemStratStore{
		lines: lines,
	}, nil
}

func FileSystemStratStoreFromFile(path string) (*FileSystemStratStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s, %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemStratStore(db)

	if err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("problem creating file system strat store, %v", err)
	}

	return store, closeFunc, nil
}

func (f *FileSystemStratStore) GetRandomStrat() string {
	return GetRandomStrat(f.lines)
}

func (f *FileSystemStratStore) ShuffleStratPile() {
	ShuffleStratPile(f.lines)
}

func (f *FileSystemStratStore) DrawTopStrat() string {
	top := DrawTopStrat(f.lines)
	return top
}
