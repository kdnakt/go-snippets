package main_test

import (
	"reflect"
	"testing"

	csv "github.com/kdnakt/go-snippets/testing-codes/csv"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		name    string
		csvPath string
		userID  string
		want    csv.User
	}{
		{
			name:    "user3 exists",
			csvPath: "./testdata/test.golden",
			userID:  "user3",
			want: csv.User{
				ID:   "user3",
				Name: "User-3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := csv.GetUser(tt.csvPath, tt.userID)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("GetUser(%s, %s) = %v, want %v", tt.csvPath, tt.userID, got, tt.want)
			}
		})
	}
}
