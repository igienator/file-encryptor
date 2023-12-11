package controllers

import (
	"path/filepath"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		filename string
		output   string
		keyfile  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful encryption",
			args: args{
				filename: "plain.txt",
				output:   "encrypted",
				keyfile:  "key.txt",
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful encryption",
			args: args{
				filename: "dupa1234.txt",
				output:   "blah",
				keyfile:  "key.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := args{
				filename: filepath.Join("testfiles", tt.args.filename),
				output:   filepath.Join("testfiles", tt.args.output),
				keyfile:  filepath.Join("testfiles", tt.args.keyfile),
			}
			if err := Encrypt(args.filename, args.output, args.keyfile); (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
