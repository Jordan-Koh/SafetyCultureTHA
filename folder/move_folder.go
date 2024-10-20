package folder

import (
	"fmt"
)

func (f *driver) MoveFolderErrorHandler(name string, dst string) error {
	if !f.CheckFolderExists(name) {
		return fmt.Errorf("Error: Source path does not exist")
	} else if !f.CheckFolderExists(dst) {
		return fmt.Errorf("Error: Destination path does not exist")
	} else if !f.CheckTargetIsNotChild(name, dst) {
		return fmt.Errorf("Error: Cannot move a folder to a child of itself")
	} else if name == dst {
		return fmt.Errorf("Error: Cannot move a folder to itself")
	} else if f.GetFolderFromName(name).OrgId != f.GetFolderFromName(dst).OrgId {
		return fmt.Errorf("Error: Cannot move a folder to a different organization")
	}
	return nil
}

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	err := f.MoveFolderErrorHandler(name, dst)
	if err != nil {
		return []Folder{}, err
	}

	return []Folder{}, nil
}
