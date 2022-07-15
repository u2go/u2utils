package u2utils

import "testing"

func TestSHA256(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hello", args{"hello"}, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.args.data); got != tt.want {
				t.Errorf("SHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"hello", args{"hello"}, "5d41402abc4b2a76b9719d911017c592", false},
		{"<empty string>", args{""}, "d41d8cd98f00b204e9800998ecf8427e", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Md5(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5() got = %v, want %v", got, tt.want)
			}
		})
	}
}
