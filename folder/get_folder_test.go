package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "test for multiple files",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "steady-insect",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect",
				},
			},
			want: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.clear-arclight",
				},
			},
		},
		{
			name:  "test for single file",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "steady-insect",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect",
				},
			},
			want: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
			},
		},
		{
			name:  "test for no occurrence",
			orgID: uuid.FromStringOrNil("no-id"),
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "steady-insect",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect",
				},
			},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name        string
		orgID       uuid.UUID
		folder_name string
		folders     []folder.Folder
		want        []folder.Folder
	}{
		{
			name:        "folder_does_not_exist",
			orgID:       uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folder_name: "fake-folder",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
		},
		{
			name:        "organisation_does_not_exist",
			orgID:       uuid.FromStringOrNil("fake-organisation"),
			folder_name: "sacred-mystique",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
		},
		{
			name:        "folder_does_not_belong_to_organisation",
			orgID:       uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folder_name: "steady-insect",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
		},
		{
			name:        "folder_exists_with_children",
			orgID:       uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
			folder_name: "upward-the-anarchist",
			folders:     folder.GetSampleData(),
			want: []folder.Folder{
				{
					Name:  "patient-prodigy",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.patient-prodigy",
				},
				{
					Name:  "obliging-microchip",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.obliging-microchip",
				},
				{
					Name:  "massive-ser-duncan",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.massive-ser-duncan",
				},
				{
					Name:  "sacred-mystique",
					OrgId: uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
					Paths: "steady-insect.helped-blackheart.many-silver-sable.upward-the-anarchist.sacred-mystique",
				},
			},
		},
		{
			name:        "folder_exists_with_no_children",
			orgID:       uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385"),
			folder_name: "sacred-mystique",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.folder_name)
			assert.Equal(t, tt.want, get)
		})
	}
}
