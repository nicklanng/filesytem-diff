package fsdiff_test

import (
	"testing"

	"github.com/nicklanng/fsdiff"
)

// func TestCompare_DifferentFilesystems(t *testing.T) {
// 	original := &fsdiff.Node{
// 		Path:        "/etc/data",
// 		Hash:        [20]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
// 		IsDirectory: true,
// 		Children: []*fsdiff.Node{
// 			&fsdiff.Node{
// 				Path: "/etc/data/a.txt",
// 				Hash: [20]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
// 			},
// 		},
// 	}
//
// 	changed := &fsdiff.Node{
// 		Path:        "/etc/otherpath",
// 		Hash:        [20]byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
// 		IsDirectory: true,
// 		Children: []*fsdiff.Node{
// 			&fsdiff.Node{
// 				Path: "/etc/otherpath/something.mp3",
// 				Hash: [20]byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
// 			},
// 		},
// 	}
//
// 	diffs := fsdiff.Compare(original, changed)
//
// 	if len(diffs) != 4 {
// 		t.Errorf("Expected 4 changes but got %d", len(diffs))
// 	}
//
// 	if diffs[0].Path != "/etc/data" {
// 		t.Error("Incorrect path in the original diff")
// 	}
// 	if diffs[0].DiffType != fsdiff.DiffTypeRemoved {
// 		t.Error("Change did not show that the original filesystem and been removed")
// 	}
//
// 	if diffs[1].Path != "/etc/data/a.txt" {
// 		t.Error("Incorrect path in the original diff")
// 	}
// 	if diffs[1].DiffType != fsdiff.DiffTypeRemoved {
// 		t.Error("Change did not show that the original filesystem and been removed")
// 	}
//
// 	if diffs[2].Path != "/etc/otherpath" {
// 		t.Error("Incorrect path in the changed diff")
// 	}
// 	if diffs[2].DiffType != fsdiff.DiffTypeAdded {
// 		t.Error("Change did not show that the changed filesystem and been removed")
// 	}
//
// 	if diffs[3].Path != "/etc/otherpath/something.mp3" {
// 		t.Error("Incorrect path in the changed diff")
// 	}
// 	if diffs[3].DiffType != fsdiff.DiffTypeAdded {
// 		t.Error("Change did not show that the changed filesystem and been removed")
// 	}
// }

func TestCompare_OneChangedFile(t *testing.T) {
	original := &fsdiff.Node{
		Path:        "/etc/data",
		Hash:        [20]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		IsDirectory: true,
		Children: []*fsdiff.Node{
			&fsdiff.Node{
				Path: "/etc/data/a.txt",
				Hash: [20]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
		},
	}

	changed := &fsdiff.Node{
		Path:        "/etc/data",
		Hash:        [20]byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		IsDirectory: true,
		Children: []*fsdiff.Node{
			&fsdiff.Node{
				Path: "/etc/data/a.txt",
				Hash: [20]byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
			},
		},
	}

	rootDiff := fsdiff.Compare(original, changed)

	if rootDiff.Path != "/etc/data" {
		t.Error("Top level node has changed")
	}

	if rootDiff.DiffType != fsdiff.DiffTypeChanged {
		t.Error("Top level directory has changed")
	}

	if len(rootDiff.Children) != 1 {
		t.Error("Should be one child")
		return
	}

	childDiff := rootDiff.Children[0]
	if childDiff.Path != "/etc/data/a.txt" {
		t.Error("Incorrect path in the diff")
	}

	if childDiff.DiffType != fsdiff.DiffTypeChanged {
		t.Error("Change did not show that the file had changed")
	}
}

func TestCompare_FolderAdded(t *testing.T) {
	original := &fsdiff.Node{
		Path:        "/etc/data",
		Hash:        [20]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		IsDirectory: true,
		Children: []*fsdiff.Node{
			&fsdiff.Node{
				Path: "/etc/data/a.txt",
				Hash: [20]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
		},
	}

	changed := &fsdiff.Node{
		Path:        "/etc/data",
		Hash:        [20]byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		IsDirectory: true,
		Children: []*fsdiff.Node{
			&fsdiff.Node{
				Path: "/etc/data/a.txt",
				Hash: [20]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
			&fsdiff.Node{
				Path:        "/etc/data/newfolder",
				Hash:        [20]byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
				IsDirectory: true,
				Children: []*fsdiff.Node{
					&fsdiff.Node{
						Path: "/etc/data/newfolder/newfile",
						Hash: [20]byte{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
					},
				},
			},
		},
	}

	rootDiff := fsdiff.Compare(original, changed)

	if rootDiff.Path != "/etc/data" {
		t.Error("Top level node has changed")
	}

	if rootDiff.DiffType != fsdiff.DiffTypeChanged {
		t.Error("Top level directory has changed")
	}

	if len(rootDiff.Children) != 2 {
		t.Error("Should be two children")
		return
	}

	childDiff := rootDiff.Children[0]
	if childDiff.Path != "/etc/data/a.txt" {
		t.Error("Incorrect path in the diff")
	}

	if childDiff.DiffType != fsdiff.DiffTypeNone {
		t.Error("Change should show that a.txt did not change")
	}

	childDiff = rootDiff.Children[1]
	if childDiff.Path != "/etc/data/newfolder" {
		t.Error("Incorrect path in the diff")
	}

	if childDiff.DiffType != fsdiff.DiffTypeAdded {
		t.Error("Folder should be shown as added")
	}

	if len(childDiff.Children) != 1 {
		t.Error("Should be one child")
		return
	}

	childDiff = childDiff.Children[0]
	if childDiff.Path != "/etc/data/newfolder/newfile" {
		t.Error("Incorrect path in the diff")
	}

	if childDiff.DiffType != fsdiff.DiffTypeAdded {
		t.Error("Change should show that newfile was added")
	}
}
