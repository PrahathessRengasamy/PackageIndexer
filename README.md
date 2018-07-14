# PackageIndexer

### Running manual:

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
### Running as a Docker container

To build the container use :

```docker build -t PackageIndexer:v1 . ```

To run the server as the container :

   ```docker run -it -p 8080:8080 PackageIndexer:v1```
   
Connections and commands are as above.

### Running tests 

In the Solution Directory run :

```go test -v```

to run the tests included and see the results

### Design Diagram

![Design](https://raw.githubusercontent.com/PrahathessRengasamy/PackageIndexer/master/PackageIndex.jpg)

### Design rationale
The design raitonale was sort of MVC - Model View and Controller.
The package index was split into the IndexModel which has the data structure map for the package index and some utility functions for the controller
The IndexController performs the operations through the utility functions provided by the IndexModel and sends result back to client.
The server is run through main.go itself accepts connections and hands them off to ManageClient (this is sort of the view aspect of MVC). ManageClient then Uses a ClientController to accept client commands
ClientController then passes them on to the MatchMaker which parses the command and matches them to the appropriate IndexController methods (Client and Index controller talk through the MatchMaker hence the name).__
 
   
