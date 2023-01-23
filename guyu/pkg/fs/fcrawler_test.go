package fs

import (
	"context"
	"testing"
)

func TestFileCrawl(t *testing.T) {
	type args struct {
		ctx    context.Context
		root   string
		vfCh   chan string
		exitCh chan bool
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	//tests = append(tests, {name: "test", args: })
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FileCrawl(tt.args.ctx, tt.args.root, tt.args.vfCh, tt.args.exitCh); (err != nil) != tt.wantErr {
				t.Errorf("FileCrawl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
