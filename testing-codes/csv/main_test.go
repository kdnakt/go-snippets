package main_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	csv "github.com/kdnakt/go-snippets/testing-codes/csv"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		name    string
		csvPath string
		userID  string
		want    csv.User
		wantErr bool
	}{
		{
			name:    "user3 exists",
			csvPath: "./testdata/test.golden",
			userID:  "user3",
			want: csv.User{
				ID:   "user3",
				Name: "User-3",
			},
			wantErr: false,
		},
		{
			name:    "no such file",
			csvPath: "./testdata/test.golden",
			userID:  "",
			want:    csv.User{},
			wantErr: true,
		},
		{
			name:    "foo doesn't exist",
			csvPath: "./testdata/nosuchfile.golden",
			userID:  "foo",
			want:    csv.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := csv.GetUser(tt.csvPath, tt.userID)
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("GetUser(%s, %s) = %v, want %v", tt.csvPath, tt.userID, got, tt.want)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	newId, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	csvPath := "./testdata/test.golden"
	user := csv.User{
		ID:   newId.String(),
		Name: "New User",
	}

	adderr := csv.AddUser(csvPath, user)
	if adderr != nil {
		t.Fatal(adderr)
	}
}
