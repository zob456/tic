package memory_handler

type Allocator struct {
	FileChunksLimit        int64
	FileChunkMemoryLimit   int64
	CompressedFileLocation string
}

func New(fileChunksLimit, fileChunkMemoryLimit int64, compressedFileLocation string) *Allocator {
	return &Allocator{
		FileChunksLimit:        fileChunksLimit,
		FileChunkMemoryLimit:   fileChunkMemoryLimit,
		CompressedFileLocation: compressedFileLocation,
	}
}

func (a *Allocator) FilterChunk() {

}
