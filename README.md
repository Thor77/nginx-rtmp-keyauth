nginx-rtmp-keyauth
==================

Very simple auth-script for [nginx-rtmp-module](https://github.com/arut/nginx-rtmp-module) using a plaintext key.

Installation
============
* Build the binary `go build server.go` or download from [Releases]()
* Install binary `mv server /usr/bin/nginx-rtmp-keyauth`
* Install systemd service `cp nginx-rtmp-keyauth.service /etc/systemd/system/`

Usage
=====
* **Optional** Modify key location by giving it to `nginx-rtmp-keyauth` as an argument
* Write your key to `/key.txt` or your custom key location
* Start the server `systemctl start nginx-rtmp-keyauth`
* Configure `http://localhost:8080/auth` as [`on_publish`](https://github.com/arut/nginx-rtmp-module/wiki/Directives#on_publish)-hook for nginx-rtmp
* Give the key as an argument when publishing `ffmpeg [...] -f flv rtmp://host:port/app/stream?key=<your key>`
