#redirector

Super simple webserver & proxy, designed for an intro Golang workshop.

This repo shouldn't be taken too seriously - it's purpose is solely to demo some features of Golang. There's no tests. No validation. No sanitisation.

##instructions

- make sure you are in the `./redirector` folder
- `go build`
- `./redirector`
- go to [http://127.0.0.1:8080/admin](http://127.0.0.1:8080/admin)

##admin interface
simply populate this with some domains - when these domains are detected, the proxy kicks in and performs the redirect. Examples of this are bundled: [http://127.0.0.1:8080/](http://127.0.0.1:8080/) which should redirect you to google
