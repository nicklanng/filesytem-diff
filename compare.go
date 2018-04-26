package fsdiff

type Diff byte

const (
	DiffNew Diff = iota
	DiffChanged
	DiffRemoved
)
