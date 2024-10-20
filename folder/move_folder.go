package folder

import (
	"fmt"
)

// error handler for GetAllChildFolders method
// Time Complexity: O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison
func (f *driver) MoveFolderErrorHandler(name string, dst string) error {
	if !f.CheckFolderExists(name) {
		return fmt.Errorf("Error: Source path does not exist")
	} else if !f.CheckFolderExists(dst) {
		return fmt.Errorf("Error: Destination path does not exist")
	} else if !f.CheckTargetIsNotChild(name, dst) { // O(N*M)
		return fmt.Errorf("Error: Cannot move a folder to a child of itself")
	} else if name == dst {
		return fmt.Errorf("Error: Cannot move a folder to itself")
	} else if f.GetFolderFromName(name).OrgId != f.GetFolderFromName(dst).OrgId {
		return fmt.Errorf("Error: Cannot move a folder to a different organization")
	}
	return nil
}

// first, identify the path to the existing source
// then, identify the path to the new destination
// iteratively check Paths of all Folder in the driver f
// if the leading strings in the path matches the existing source,
// then create a new folder concatenating the new path + old path less the path to source
// otherwise include the existing folder as is
//
// Time Complexity: O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	err := f.MoveFolderErrorHandler(name, dst)
	if err != nil {
		return []Folder{}, err
	}

	folders := f.folders

	old_parent := f.GetFolderFromName(name)
	old_parent_path := old_parent.Paths[:len(old_parent.Paths)-len(old_parent.Name)-1]

	new_parent := f.GetFolderFromName(dst)
	new_parent_path := new_parent.Paths

	res := []Folder{}

	for _, f := range folders {
		// check if Folder f starts with the the old source's parent path
		if (len(f.Paths) > len(old_parent_path)) && (f.Paths[:len(old_parent_path)] == old_parent_path) {
			fmt.Println("hit")
			fmt.Println(f.Paths)
			res = append(res, Folder{
				Name:  f.Name,
				OrgId: f.OrgId,
				Paths: new_parent_path + f.Paths[len(old_parent_path):],
			})
		} else {
			fmt.Println("nohit")
			res = append(res, f)
		}
	}

	return res, nil
}
