# PackageIndexer

###Running manual:

main.go is the server that manages client requests

In a new terminal open the ```/Solution   ```  directory and run

```go build``` 

Generates the ```Solution``` binary  then run:

```./Solution``` 

to start the server.

To connect to it and send commands as client in a new terminal run netcat:

```nc localhost 8080```

Commands from clients follow this pattern:

```
<command>|<package>|<dependencies>\n
```

Where:
* `<command>` is mandatory, and is either `INDEX`, `REMOVE`, or `QUERY`
* `<package>` is mandatory, the name of the package referred to by the command, e.g. `mysql`, `openssl`, `pkg-config`, `postgresql`, etc.
* `<dependencies>` is optional, and if present it will be a comma-delimited list of packages that need to be present before `<package>` is installed. e.g. `cmake,sphinx-doc,xz`
* The message always ends with the character `\n`

Here are some sample messages:
```
INDEX|cloog|gmp,isl,pkg-config\n
INDEX|ceylon|\n
REMOVE|cloog|\n
QUERY|cloog|\n
```
###Running Docker

:
