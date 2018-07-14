package main

import (
	"testing"
	"net"
)

//Tests the ManageClient function



func TestManageClient(t *testing.T) {
	link, err := net.Listen("tcp", "")
	if err != nil {
		return
	}
	defer link.Close()

	type args struct {
		ln net.Listener
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test client connection bad listener.", args: args{link}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go ManageClient(tt.args.ln)
		})
	}
}

