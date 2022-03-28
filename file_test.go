package u2utils

import (
	"path/filepath"
	"testing"
)

func TestFileExistsParent(t *testing.T) {
	type args struct {
		dir      string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"t1.txt", args{"mock/d1", "t1.txt"}, "mock/d1/t1.txt", false},
		{"mock", args{"mock/d1", "mock"}, "mock", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileExistsParent(tt.args.dir, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileExistsParent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want, err := filepath.Abs(tt.want)
			if err != nil {
				t.Errorf("FileExistsParent() error = %v, wantErr %v", err, tt.want)
			}
			if got != want {
				t.Errorf("FileExistsParent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
