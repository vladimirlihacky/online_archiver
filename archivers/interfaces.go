package archivers

/**
* Any archiver implementation shall fulfill this interface
* Can be used to implement ZIP/RAR/TAR/7z etc.
 */
type Archiver interface {
	compress(data []byte) ([]byte, error)
	decompress(data []byte) ([]byte, error)
}
