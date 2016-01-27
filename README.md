# gwebimg
Create a local web server to serve a local directory as an web image gallery with swipe. 
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
Running command `gwebimg -h=127.0.0.1 -i=top` will serve directory top as 02 chapter [dir1 & dir2], each chapter is a swipe gallery of file1, file2 and file3, file4.

This code use pure css (<http://purecss.io>), zeptojs (<http://zeptojs.com>) and mobify scooch <https://github.com/mobify/scooch>
