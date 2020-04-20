
set GOOS=linux
set GOARCH=arm
set GOARM=7
set CGO=0

go build -o hello -ldflags="-s -w" hello.go
: 1.2 MB
scp hello root@192.168.1.1:

