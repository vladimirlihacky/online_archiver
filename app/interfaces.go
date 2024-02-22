package app

import "online_archiver/archivers"

type App interface {
	Start() error
	Compress(archiver archivers.Archiver, data []byte)
	Decompress(archiver archivers.Archiver, data []byte)
}
