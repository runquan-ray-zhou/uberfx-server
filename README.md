# Multi App Central Backend Server

This central backend server will serve as the backbone to all my current and future apps to handle calling external APIs, data storage, logging, auth and so forth. Built with uber's Fx application framework in Go.

Currently serving the following apps.

- [https://linknyc-finder.netlify.app/](https://linknyc-finder.netlify.app/)
- [https://pocket-dictionary-app.netlify.app/](https://pocket-dictionary-app.netlify.app/)
- [https://quiz-me-trivia-app.netlify.app/](https://quiz-me-trivia-app.netlify.app/)

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

### Health Check

```
$ curl -X GET http://localhost:8080/
```

### LinkNYC URL route pattern

```
$ curl -X GET http://localhost:8080/linknyc
```

### Pocket Dictionary URL route pattern

```
$ curl -X GET http://localhost:8080/pocketdictionary
```

### Quiz Me URL route pattern

```
$ curl -X GET http://localhost:8080/quizme
```
