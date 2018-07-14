package PackageIndexer


import (
	"net"
)
//A separate controller to keep the operations on the Package Index and respond to client
//using the same names as the user commands for the funcions
// Index updates the package

func Index(conn net.Conn, p Package, resp Response) bool {
	indexed := false
	if len(p.Dependencies) == 0 || HasDependencies(p) {
		go AddPackage(p.Name, p.Dependencies)
		conn.Write([]byte(resp.StatusOk + "\n"))
		indexed = true
	} else if IsPresent(p) {
		conn.Write([]byte(resp.StatusOk + "\n"))
		indexed = true
	} else {
		conn.Write([]byte(resp.StatusFail + "\n"))
	}
	return indexed
}

// Remove removes a package from the package index

func Remove(conn net.Conn, p Package, resp Response) bool {
	removed := false
	if !TransitiveDep(p) || !IsPresent(p) {
		go DeletePackage(p.Name)
		conn.Write([]byte(resp.StatusOk + "\n"))
		removed = true
	} else {
		conn.Write([]byte(resp.StatusFail + "\n"))
		removed = false
	}
	return removed
}

// Query searches for a given package in the package index

func Query(conn net.Conn, p Package, resp Response) bool {
	queried := false
	if IsPresent(p) {
		conn.Write([]byte(resp.StatusOk + "\n"))
		queried = true
	} else {
		conn.Write([]byte(resp.StatusFail + "\n"))
	}
	return queried
}
