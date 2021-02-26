# transformurls
A simple tool that transforms a url or a list of urls to multiple variations of urls by path. 

## Goal

```
╰─$ cat examples.txt
https://www.example.com/abc/that/auf.php
https://www.example.com/abc/thatsl/xsauf.php

╰─$ cat examples.txt | go run main.go
https://www.example.com/
https://www.example.com/abc/
https://www.example.com/abc/that/
https://www.example.com/
https://www.example.com/abc/
https://www.example.com/abc/thatsl/
```

## Install

`go get -u github.com/Shaked/transformurls` 
