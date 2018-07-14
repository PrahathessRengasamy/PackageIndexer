package main
import (
	"net"
	"testing"
)

func TestClientController(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()
	//Defining struct to make inserting test cases easier
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test connection.", args: args{client}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go ClientController(tt.args.conn)
		})
	}
}
func TestMatchMaker(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()

	response := Response{"OK", "FAIL", "ERROR"}
	command := Command{"INDEX", "REMOVE", "QUERY"}
	pkg1 := Package{Name: "coolpkg", Dependencies: []string{}}
	pkg2 := Package{}
//Defining struct to make inserting test cases easier
	type args struct {
		conn net.Conn
		resp Response
		comm Command
		cmd  string
		pkg  Package
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Invalid Command",
			args: args{
				client,
				response,
				command,
				"",
				pkg1}},
		{
			name: "Test Empty Package",
			args: args{
				client,
				response,
				command,
				command.Index,
				pkg2}},
		{
			name: "Test Index",
			args: args{
				client,
				response,
				command,
				command.Index,
				pkg1}},
		{
			name: "Test Remove",
			args: args{
				client,
				response,
				command,
				command.Remove,
				pkg1}},
		{
			name: "Test Query",
			args: args{
				client,
				response,
				command,
				command.Query,
				pkg1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go MatchMaker(
				tt.args.conn,
				tt.args.resp,
				tt.args.comm,
				tt.args.cmd,
				tt.args.pkg)
		})
	}
}
