package main



import (
"net"
"testing"
)

func TestHasDependencies(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()
	pkg1 := Package{Name: "X", Dependencies: []string{}}
	pkg2 := Package{Name: "Z", Dependencies: []string{"X"}}

	type args struct {
		pkg Package
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Dependency Doesn't Exist",
			args: args{pkg1},
			want: false},
		{
			name: "Test Dependency Does Exist",
			args: args{pkg2},
			want: true},
	}
	package_index[pkg1.Name] = []string{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDependencies(tt.args.pkg); got != tt.want {
				t.Errorf("HasDependencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPresent(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()
	pkg1 := Package{Name: "A", Dependencies: []string{}}
	pkg2 := Package{Name: "B", Dependencies: []string{}}

	type args struct {
		pkg Package
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Doesn't Exist",
			args: args{pkg1},
			want: false},
		{
			name: "Test Exists",
			args: args{pkg2},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			package_index[pkg2.Name] = []string{}
			if got := IsPresent(tt.args.pkg); got != tt.want {
				t.Errorf("IsPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransitive(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()
	pkg1 := Package{Name: "vimgo", Dependencies: []string{"vim", "nvim"}}
	pkg2 := Package{Name: "nvim", Dependencies: []string{}}

	type args struct {
		pkg Package
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Dependency Doesn't Exist",
			args: args{pkg1},
			want: false},
		{
			name: "Test Dependency Does Exist",
			args: args{pkg2},
			want: true},
	}
	package_index[pkg1.Name] = pkg1.Dependencies
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransitiveDep(tt.args.pkg); got != tt.want {
				t.Errorf("Transitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
