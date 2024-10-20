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

	// example: feel free to change the data structure, if slice is not what you want
	folders []Folder
}

func NewDriver(folders []Folder) IDriver {
	return &driver{
		// initialize attributes here
		folders: folders,
	}
}

func (f *driver) GetFolderFromName(name string) Folder {
	folders := f.folders

	for _, f := range folders {
		directories := strings.Split(f.Paths, ".")
		if slices.Contains(directories, name) {
			return f
		}
	}

	return Folder{}
}

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

func (f *driver) CheckTargetIsNotChild(name string, dst string) bool {
	target_folder := f.GetFolderFromName(dst)

	return !IsParentOf(name, target_folder)
}

func (f *driver) CheckFolderExists(name string) bool {
	folders := f.folders

	for _, f := range folders {
		if f.Name == name {
			return true
		}
	}
	return false
}

func (f *driver) CheckOrgExists(orgID uuid.UUID) bool {
	folders := f.folders

	for _, f := range folders {
		if f.OrgId == orgID {
			return true
		}
	}
	return false
}
