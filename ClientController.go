package PackageIndexer

import (
	"bufio"
	"log"
	"net"
	"bytes"
)

// Response statuses

type Response struct {
	StatusOk    string
	StatusFail  string
	StatusError string
}

// Valid Input commands
type Command struct {
	Index  string
	Remove string
	Query  string
}

// ClientController client connection and message Controller

func ClientController(conn net.Conn) {
	resp := Response{"OK", "FAIL", "ERROR"}
	command := Command{"INDEX", "REMOVE", "QUERY"}

	defer func() {
		conn.Close()
	}()

	br := bufio.NewReader(conn)
	for {
		// Read in tokens, delimited by newline.
		b, err := br.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			return
		}
		cmd, pkg, deps := Parser(b)
		p := Package{Name: pkg, Dependencies: deps}
		go MatchMaker(conn, resp, command, cmd, p)
	}
}

// MatchMaker passes the package along to the necessary response function.
//
func MatchMaker(conn net.Conn, resp Response, comm Command, cmd string, pkg Package) {
	if pkg.Name == "" ||
		cmd != comm.Index &&
			cmd != comm.Remove &&
			cmd != comm.Query {
		conn.Write([]byte(resp.StatusError + "\n"))
	} else if cmd == comm.Index {
		Index(conn, pkg, resp)
	} else if cmd == comm.Remove {
		Remove(conn, pkg, resp)
	} else if cmd == comm.Query {
		Query(conn, pkg, resp)
	} else {
		conn.Write([]byte(resp.StatusError + "\n"))
	}
}
// Parsers <command>|<package>|<dependencies>\n into proper commands
func Parser(ba []byte) (string, string, []string) {
	buf := new(bytes.Buffer)
	cmd := ""
	pkg := ""
	deps := []string{}
	i := 0

	for b := range ba {
		// 124 is the decimal value for ASCII pipe.
		if ba[b] != 124 {
			buf.WriteString(string(ba[b]))
		} else {
			if i == 0 {
				cmd = buf.String()
				i++
			} else if i == 1 {
				pkg = buf.String()
				i++
			} else {
				deps = append(deps, buf.String())
			}
			buf.Reset()
		}
	}
	return cmd, pkg, deps
}
