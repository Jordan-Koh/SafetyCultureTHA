package folder

import (
	"fmt"

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
// solution has a complexity of O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison complexity
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	err := f.GetAllChildFoldersErrorHandler(orgID, name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// gets all folders that has OrgID = orgID
	folders := f.GetFoldersByOrgID(orgID)

	res := []Folder{}

	// iterates through every folder which belongs to orgID
	for _, f := range folders {
		// splits directory string into individual folder name tokens until a folder matches name
		if IsParentOf(name, f) {
			res = append(res, f)
		}
	}

	return res
}
