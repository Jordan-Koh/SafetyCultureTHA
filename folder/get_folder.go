package folder

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

// returns Folder object from its name
//
// Time Complexity: O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison
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

// error handler for GetAllChildFolders method
//
// Time Complexity: O(N+N+N) -> O(N) where
// N = number of folders
// assuming O(1) string comparison
func (f *driver) GetAllChildFoldersErrorHandler(orgID uuid.UUID, name string) error {
	if !f.CheckFolderExists(name) {
		return fmt.Errorf("Error: Folder does not exist")
	} else if !f.CheckOrgExists(orgID) {
		return fmt.Errorf("Error: Organization does not exist")
	} else if f.GetFolderFromName(name).OrgId != orgID {
		return fmt.Errorf("Error: Folder does not exist in the specified organization")
	}
	return nil
}

// gets child folders by iteratively checking if name exists in one of the parent directories
//
// Time Complexity: O(N+N+(N*M)) -> O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	err := f.GetAllChildFoldersErrorHandler(orgID, name)
	if err != nil {
		fmt.Println(err)
		return []Folder{}
	}

	// gets all folders that has OrgID = orgID
	folders := f.GetFoldersByOrgID(orgID)

	res := []Folder{}

	// iterates through every folder which belongs to orgID
	for _, f := range folders {
		// checks if name is a parent of current folder f
		if IsParentOf(name, f) {
			res = append(res, f)
		}
	}

	return res
}
