# gong
Simple HTTP server in Go.

Running the application starts a HTTP server in port 8080 serving the contents of the current working directory. If index.html is present, root serves that. Else it serves directory listing.

## Arguments
+ -port *port*

  Changes the server port
  
+ -root *path*

  Changes the root directory to serve
