# gwebimg
Single binary to create a local web server to serve a local directory as an web image gallery with touch support. 
Example: 
Local dir:
```
top/
   dir1/
        file1.jpg
        file2.jpg
   dir2/
        file3.jpg
        file4.jpg
```
Running command `gwebimg -h=127.0.0.1 -p=8080 -i=top` will serve directory top as 02 chapter [dir1 & dir2], each chapter is a swipe gallery of file1, file2 and file3, file4. Use web browser to access http://localhost:8080/ to view. Basic keyboard navigation using left, right arrow key and page up & page down key is also supported.

This code use gorilla toolkit (<http://www.gorillatoolkit.org>), pure css (<http://purecss.io>), zeptojs (<http://zeptojs.com>) and mobify scooch <https://github.com/mobify/scooch>

How to compile:
- set GOPATH
- go get github.com/gorilla/mux
- go get github.com/gorilla/context
- go get github.com/gorilla/handlers
- go build

Set GOARCH & GOOS for cross compiling to other arch.