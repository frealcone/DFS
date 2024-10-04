package master_server

import "fmt"

// FileSystem manages all files stored in DFS.
type FileSystem interface {
	// Touch creates a new file.
	Touch(string) error
	// GetHandle finds a chunk handle based on filename and offset.
	GetHandle(string, uint64) (ChunkHandle, error)
}

// F2Hs holds <filename, chunk handles> pairs
type F2Hs map[string][]ChunkHandle

func (f *F2Hs) Touch(filename string) error {
	if _, ok := (*f)[filename]; ok {
		return fmt.Errorf("touch: file exists: %s", filename)
	}
	(*f)[filename] = []ChunkHandle{}
	return nil
}

func (f *F2Hs) GetHandle(filename string, offset uint64) (ChunkHandle, error) {
	handles, ok := (*f)[filename]
	if !ok {
		return ChunkHandle{}, fmt.Errorf("%s: no such file", filename)
	}

	idx := offset / ChunkSize
	if idx >= uint64(len(handles)) {
		return ChunkHandle{}, fmt.Errorf("offset too large: %d", offset)
	}

	return handles[idx], nil
}
