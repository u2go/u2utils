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
