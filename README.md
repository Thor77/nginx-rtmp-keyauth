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
* Start the server `systemctl start nginx-rtmp-keyauth`
* Configure `http://localhost:8080` as [`on_publish`](https://github.com/arut/nginx-rtmp-module/wiki/Directives#on_publish)-hook for nginx-rtmp
