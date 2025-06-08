# Multi App Central Backend Server

This central backend server will serve as the backbone to all my current and future apps to handle calling external APIs, data storage, logging, auth and so forth. Built with uber's Fx application framework in Go.

### Set Up

```
$ go mod init github.com/runquan-ray-zhou/uberfx-server
```

```
$ go mod vendor
```

```
$ go mod tidy
```

### Start Server

```
$ go run main.go
```
