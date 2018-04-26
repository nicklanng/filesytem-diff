package fsdiff

import (
	"crypto/sha1"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrPathNotFound          = errors.New("path not found")
	ErrRootMustBeDirectory   = errors.New("root must be a directory")
	ErrFailedToReadDirectory = errors.New("failed to read directory")
	ErrFailedToReadFile      = errors.New("failed to read file")
	ErrFailedToComputeHash   = errors.New("failed to compute hash")
)

type FileSystem struct {
	Path string
	Root *Directory
}

type Directory struct {
	Name        string
	Hash        [20]byte
	Directories []*Directory
	Files       []*File
}

type File struct {
	Name string
	Hash [20]byte
}

func Build(root string) (*FileSystem, error) {
	stat, err := os.Stat(root)
	if err != nil {
		return nil, ErrPathNotFound
	}

	if !stat.IsDir() {
		return nil, ErrRootMustBeDirectory
	}

	rootDirectory, err := walkDirectory(root)
	if err != nil {
		return nil, err
	}

	return &FileSystem{
		Path: root,
		Root: rootDirectory,
	}, nil
}

func walkDirectory(path string) (*Directory, error) {
	// find files
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, ErrFailedToReadDirectory
	}

	// look through all files and folders
	files := []*File{}
	directories := []*Directory{}
	for _, fileInfo := range fileInfos {
		fullFilePath := filepath.Join(path, fileInfo.Name())

		// is file a directory
		stat, err := os.Stat(fullFilePath)
		if err != nil {
			return nil, ErrPathNotFound
		}

		if stat.IsDir() {
			directory, err := walkDirectory(fullFilePath)
			if err != nil {
				return nil, err
			}

			directories = append(directories, directory)
		} else {
			file, err := walkFile(fullFilePath)
			if err != nil {
				return nil, err
			}

			files = append(files, file)
		}
	}

	// calculate folder hash
	directoryHash := sha1.New()
	for _, f := range files {
		if _, err := directoryHash.Write(f.Hash[:]); err != nil {
			return nil, ErrFailedToComputeHash
		}
	}
	for _, d := range directories {
		if _, err := directoryHash.Write(d.Hash[:]); err != nil {
			return nil, ErrFailedToComputeHash
		}
	}
	var hash [20]byte
	copy(hash[:19], directoryHash.Sum(nil))

	return &Directory{
		Name:        filepath.Base(path),
		Hash:        hash,
		Directories: directories,
		Files:       files,
	}, nil
}

func walkFile(path string) (*File, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, ErrFailedToReadFile
	}

	return &File{
		Name: filepath.Base(path),
		Hash: sha1.Sum(file),
	}, nil
}
