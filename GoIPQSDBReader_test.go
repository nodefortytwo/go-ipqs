package ipqs

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {

	ipv4, _ := os.Open("test-files/ipv4.ipqs")

	f, _ := os.Open("test-files/ipv4.ipqs")
	buff, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	type args struct {
		r Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"file", args{r: ipv4}, false},
		{"file", args{r: bytes.NewReader(buff)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Open(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
