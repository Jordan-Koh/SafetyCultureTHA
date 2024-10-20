package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
)

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

// import (
// 	"fmt"

// 	"github.com/georgechieng-sc/interns-2022/folder"

// 	"github.com/gofrs/uuid"
// )

// func main() {

// 	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

// 	res := folder.GetAllFolders()

// 	folderDriver := folder.NewDriver(res)
// 	childFolders := folderDriver.GetAllChildFolders(orgID, "magnetic-sinister-six")

// 	folder.PrettyPrint(childFolders)
// }

func main() {
	res := folder.GetAllFolders()

	folderDriver := folder.NewDriver(res)
	movedFolders, err := folderDriver.MoveFolder("magnetic-sinister-six", "smashing-abyss")

	if err != nil {
		fmt.Println(err)
	}

	folder.PrettyPrint(movedFolders)
}
