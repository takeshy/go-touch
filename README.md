# go-touch - inotify proxy

## Abstract

go-touch is inotify proxy written in Golang. go-touch listens changed file name with tcp connection, and touch that file.
I wrote go-touch-client program for Vim and VisualStudio Code.

## Use cases

You can use go-touch to Docker for Windows on shared directory.
because currently, inotify does not work on Docker for Windows. 
execute go-touch on docker container and if a file is changed, editor(Vim or VisualStudio Code) send that file name to go-touch then go-touch touch that file.  So App as Rails, webpack-dev-server etc. can detect file changes.

## Installation

### Binary for Linux X64

https://s3-ap-northeast-1.amazonaws.com/takesy-work/go-touch

### From source codes

```sh
$ go get github.com/takeshy/go-touch
```
```sh
$ wget https://raw.githubusercontent.com/takeshy/go-touch/master/main.go
$ go build
```

### For Vim

ex.  go-touch execute on $HOME and shared C:\Users\USERNAME\work\PROJECT => $HOME/PROJECT
add below your .vimrc
```.vimrc
function! SendSavedFile() abort
  let removePath = "/work"
  let appendPath = ""
  let port = 7650
  if match(expand('%:p'), expand('~') . removePath) == 0
    let filePath = substitute(expand('%:p'), expand('~') . removePath, appendPath, '')
    let channel = ch_open('localhost:' . port)
    call ch_sendraw(channel, filePath . "\n")
  endif
endfunction
autocmd BufWritePost * :call SendSavedFile()
```
change variables value if you need
- removePath ...  remove path continuous $USERPROFILE <br/>
**ex. /work if /Users/USERNME/work/PROJECT -> /home/USERNAME/PROJECT**
- appendPath ... append path continuous $HOME <br/>
**ex. /app if /Users/USERNAME/PROJECT -> /home/USERNAME/app/PROJECT**
- port ... liten port go-touch

### For Visual Studio Code

install extenction go-touch-client
configuration
- "go-touch-client.removePath": remove path continuous $USERPROFILE  default "" <br/>
**ex. /work if /Users/USERNME/work/PROJECT -> /home/USERNAME/PROJECT**
- "go-touch-client.appendPath": append path continuous $HOME default "" <br/>
**ex. /app if /Users/USERNAME/PROJECT -> /home/USERNAME/app/PROJECT**
- "go-touch-client.port": go-touch port default 7650

## Execution

On your $HOME directory execute the following command:

```sh
$ go-touch &
```
-p port for listen default: 7650
-h ip for listen default 0.0.0.0

you need shared port when execute docker-run 
ex. 
```
docker run -p 7650:7650 -v ~/work/PROJECT:/home/hogohoge/PROJECT IMAGENAME START_PROG
```
## Licence
MIT
