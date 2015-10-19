# gong
Simple HTTP server application in Go.

Running the application starts a HTTP server in port 8080 serving the contents of the current working directory. If index.html is present, root serves that. Else it serves directory listing.

## Arguments
+ -port *port*

  Changes the server port
  
+ -root *path*

  Changes the root directory to serve

## Installation
`go get github.com/ComSecNinja/gong`
`go install github.com/ComSecNinja/gong`

## Add Windows context menu item
Being able to right-click and choose "Serve these files" on Windows' Explorer is handy. This is how you achieve it:
* Press `WIN+R`, type in `regedit.exe` and press Enter
* If you want to add this feature for all users, expand `HKEY_CLASSES_ROOT`. Else expand `HKEY_CURRENT_USER`.
* Navigate to `SOFTWARE\Classes\Directory`
* To add the feature for right-clicking folder background `Background\shell`. If for right-clicking folder, expand `shell`.
* Right-click `shell` you've expanded, choose `New > Key` and enter e.g. *Serve these files*.
* Right click the key you created, choose `New > key` and enter *command* (no choice here).
* Select `command` and on the right hand side right-click the `(Default)` and click `Modify...`
* Type in the absolute path of *gong.exe*. It should be in `%GOPATH%/bin/gong.exe`.
* If you want a different port for this function, append ` -port=` and the port number you desire.
* If you're in `Directory\shell` you **MUST** append ` -root=%1` for this to work properly.
* After clicking `OK` you should now be able to use gong this way.
