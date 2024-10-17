package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()
	fld := "healthy-deathstrike"
	fld2 := "equipped-hypno-hustler"
	// example usage
	folderDriver := folder.NewDriver(res)
	fmt.Printf("Folders for orgID: %s\n", orgID)
	//orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	//folder.PrettyPrint(orgFolder)
	//fmt.Printf("\nFolders for Parent: %s\n", fld)
	childFolder := folderDriver.GetAllChildFolders(orgID, fld)
	folder.PrettyPrint(childFolder)
	fmt.Println()
	folderDriver.MoveFolder(fld, fld2)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	folder.PrettyPrint(orgFolder)
	
}
