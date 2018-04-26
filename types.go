package fsdiff

type FileSystem struct {
	Path string
	Root *Node
}

type Node struct {
	Path        string
	Hash        [20]byte
	IsDirectory bool
	Children    []*Node
}
