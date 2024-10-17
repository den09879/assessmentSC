package folder

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders
	res := []Folder{}
	path := ""
	children := []Folder{}
	for i := range folders {
		if folders[i].Name == dst {
			path = folders[i].Paths + "." + name
			break
		}
	}

	for i := range folders {

		if folders[i].Name == name {
			children = f.GetAllChildFolders(folders[i].OrgId, folders[i].Name)
			folders[i].Paths = path
		}
		for j := range children {

			if children[j].Name == folders[i].Name {
				folders[i].Paths = path + "." + folders[i].Name
				print(path)
			}
		}
	}
	return res, nil
}
