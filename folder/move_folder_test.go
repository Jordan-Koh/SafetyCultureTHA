package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name        string
		source_name string
		dest_name   string
		folders     []folder.Folder
		want        []folder.Folder
		want_err    bool
	}{
		{
			name:        "source_path_does_not_exist",
			source_name: "fake_source_name",
			dest_name:   "creative-scalphunter",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
			want_err:    true,
		},
		{
			name:        "destination_path_does_not_exist",
			source_name: "creative-scalphunter",
			dest_name:   "fake_dest_name",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
			want_err:    true,
		},
		{
			name:        "attempt_to_move_folder_into_child",
			source_name: "creative-scalphunter",
			dest_name:   "clear-arclight",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
			want_err:    true,
		},
		{
			name:        "attempt_to_move_folder_to_itself",
			source_name: "clear-arclight",
			dest_name:   "clear-arclight",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
			want_err:    true,
		},
		{
			name:        "attempt_to_move_folder_to_other_organisation",
			source_name: "creative-scalphunter",
			dest_name:   "noble-vixen",
			folders:     folder.GetSampleData(),
			want:        []folder.Folder{},
			want_err:    true,
		},
		{
			name:        "move_folder_to_parent",
			source_name: "magnetic-sinister-six",
			dest_name:   "noble-vixen",
			folders: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "nearby-secret",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name:  "magnetic-sinister-six",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six",
				},
				{
					Name:  "stirred-rainbow",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six.stirred-rainbow",
				},
			},
			want: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "nearby-secret",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name:  "magnetic-sinister-six",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.magnetic-sinister-six",
				},
				{
					Name:  "stirred-rainbow",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.magnetic-sinister-six.stirred-rainbow",
				},
			},
			want_err: false,
		},
		{
			name:        "move_folder_to_other_parent",
			source_name: "magnetic-sinister-six",
			dest_name:   "stunning-horridus",
			folders: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "nearby-secret",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name:  "magnetic-sinister-six",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six",
				},
				{
					Name:  "stirred-rainbow",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret.magnetic-sinister-six.stirred-rainbow",
				},
				{
					Name:  "stunning-horridus",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stunning-horridus",
				},
			},
			want: []folder.Folder{
				{
					Name:  "noble-vixen",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen",
				},
				{
					Name:  "nearby-secret",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "noble-vixen.nearby-secret",
				},
				{
					Name:  "magnetic-sinister-six",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stunning-horridus.magnetic-sinister-six",
				},
				{
					Name:  "stirred-rainbow",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stunning-horridus.magnetic-sinister-six.stirred-rainbow",
				},
				{
					Name:  "stunning-horridus",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stunning-horridus",
				},
			},
			want_err: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.source_name, tt.dest_name)
			if tt.want_err {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, get)
		})
	}
}
