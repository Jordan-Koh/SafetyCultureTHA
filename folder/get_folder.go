package folder

import (
	"fmt"
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

// this error check assumes all folders has its own folder entry in the data
// ie. if alpha.beta exists in f.folders then alpha must also exist in f.folders
func (f *driver) FolderErrorHandler(orgID uuid.UUID, name string) error {
	folders := f.folders

	orgid_exist := false         // checks if orgID exists in data
	path_exist := false          // checks if folder exists in data
	folder_in_org_exist := false // checks if folder belongs to organisation orgID

	for _, f := range folders {

		// orgID exists
		if f.OrgId == orgID {
			orgid_exist = true
		}

		// folder exists
		if f.Name == name {
			path_exist = true
		}

		// folder belongs to orgID
		if (f.OrgId == orgID) && (f.Name == name) {
			folder_in_org_exist = true
			break
		}
	}

	if !orgid_exist {
		return fmt.Errorf("orgID does not exist")
	} else if !path_exist {
		return fmt.Errorf("folder does not exist")
	} else if !folder_in_org_exist {
		return fmt.Errorf("folder does not exist in the specified organization")
	} else {
		return nil
	}
}

// gets child folders by iteratively checking if name exists in one of the parent directories
// solution has a complexity of O(N*M) where
// N = number of folders
// M = largest number of parents a folder has
// assuming O(1) string comparison complexity
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	err := f.FolderErrorHandler(orgID, name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(name)
	// gets all folders that has OrgID = orgID
	folders := f.GetFoldersByOrgID(orgID)

	res := []Folder{}

	// iterates through every folder which belongs to orgID
	for _, f := range folders {
		// splits directory string into individual folder name tokens until a folder matches name
		for _, path := range strings.Split(f.Paths, ".") {
			if (path == name) && (f.Name != name) {
				res = append(res, f)
				break
			}
		}
	}

	return res
}
