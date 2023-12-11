package controllers

import (
	"path/filepath"
	"testing"
)

func TestDecrypt(t *testing.T) {
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
			name: "Successful decryption",
			args: args{
				filename: "encrypted_2",
				output:   "decrypted.txt",
				keyfile:  "key.txt",
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful decryption",
			args: args{
				filename: "dupa1234",
				output:   "blah.txt",
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
			if err := Decrypt(args.filename, args.output, args.keyfile); (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
