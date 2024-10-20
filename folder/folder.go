package folder

import (
	"slices"
	"strings"

	"github.com/gofrs/uuid"
)

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// data structure was kept as a slice as changing could result in compatibility issues regarding the testing of data

	// the second option was to use a tree like structure so that operations involving moving and identifying child directories would have an improved best case complexity

	// the final option i could think of was by using a two way linked tree like structure with a hashmap for O(1) access to any node.
	// the final option would result in O(1) lookup and move operations; and O(n) for identifying all child directories of any given folder

	folders []Folder
}

func NewDriver(folders []Folder) IDriver {
	return &driver{
		// initialize attributes here
		folders: folders,
	}
}

// checks if name is a parent of Folder f
//
// Time Complexity: O(M) where
// M = Folder f's number of parents
// assuming O(1) string comparison
func IsParentOf(name string, f Folder) bool {
	parents := strings.Split(f.Paths, ".")

	if len(parents) <= 0 {
		return false
	}

	if slices.Contains(parents[:len(parents)-1], name) {
		return true
	} else {
		return false
	}
}

// checks if dst is not a child of name
//
// Time Complexity: O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison
func (f *driver) CheckTargetIsNotChild(name string, dst string) bool {
	target_folder := f.GetFolderFromName(dst)

	return !IsParentOf(name, target_folder)
}

// checks if name exists in the driver f
//
// Time Complexity: O(N) where
// N = number of folders
// assuming O(1) string comparison
func (f *driver) CheckFolderExists(name string) bool {
	folders := f.folders

	for _, f := range folders {
		if f.Name == name {
			return true
		}
	}
	return false
}

// checks if orgID exists in the driver f
//
// Time Complexity: O(N) where
// N = number of folders
// assuming O(1) UUID comparison
func (f *driver) CheckOrgExists(orgID uuid.UUID) bool {
	folders := f.folders

	for _, f := range folders {
		if f.OrgId == orgID {
			return true
		}
	}
	return false
}
