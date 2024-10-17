package folder

import (
	"errors"
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.folders
	res := []Folder{}
	length := 0
	uuid_flag := false
	for i := range folders {
		currPath := strings.Split(folders[i].Paths, ".")
		if folders[i].Name == name {
			length = len(currPath)
			if folders[i].OrgId == orgID {
				uuid_flag = true
			}

			break
		}
	}

	if length == 0 {
		fmt.Println(errors.New("Folder does not exist"))
		return res
	} else {
		if !uuid_flag {
			fmt.Println(errors.New("Folder does not exist in the specified organization"))
			return res
		}
	}

	for i := range folders {
		currPath := strings.Split(folders[i].Paths, ".")
		if uuid_flag && length < len(currPath) {
			if string(currPath[length-1]) == name {
				res = append(res, folders[i])
			}
		}

	}

	return res

}
