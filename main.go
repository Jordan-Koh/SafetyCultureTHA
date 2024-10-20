// package main

// import (
// 	"fmt"

// 	"github.com/georgechieng-sc/interns-2022/folder"
// 	"github.com/gofrs/uuid"
// )

// func main() {
// 	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

// 	res := folder.GetAllFolders()

// 	// example usage
// 	folderDriver := folder.NewDriver(res)
// 	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

// 	folder.PrettyPrint(res)
// 	fmt.Printf("\n Folders for orgID: %s", orgID)
// 	folder.PrettyPrint(orgFolder)
// }

package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"

	"github.com/gofrs/uuid"
)

func main() {
	fmt.Println("hi")
	fmt.Println(folder.IsParentOf(
		"whack",
		folder.Folder{
			Name:  "name",
			OrgId: uuid.FromStringOrNil("id"),
			Paths: "this.is.my.name",
		},
	))
	fmt.Println("bye")

	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// res := folder.GetAllFolders()

	// folderDriver := folder.NewDriver(res)
	// childFolders := folderDriver.GetAllChildFolders(orgID, "magnetic-sinister-six")

	// folder.PrettyPrint(childFolders)
}
