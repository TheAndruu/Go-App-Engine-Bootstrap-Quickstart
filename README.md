# Go, App Engine, and Bootstrap templates
Quickstart project for Google App Engine with Go SDK and Bootstrap themed templates.

# Overview

More details and suggested use are available in the original article here: https://www.m0d3rnc0ad.com/post/go-gae-bootstrap-quickstart/

# What it looks like

The demo running at [http://bootstrap.m0d3rnc0ad.appspot.com/](http://bootstrap.m0d3rnc0ad.appspot.com/), demonstrates the ability to render templated pages inheriting all the Bootstrap components

# Getting the code

In addition to checking out this git repo, you can fetch the same quickstart code with ```go get``` if you already have the App Engine Go SDK installed (instructions: https://www.m0d3rnc0ad.com/post/start-on-appengine/) with the following:

```bash
go get github.com/TheAndruu/Go-App-Engine-Bootstrap-Quickstart
cd $GOPATH/src/github.com/TheAndruu/Go-App-Engine-Bootstrap-Quickstart
goapp serve
```

Navigating a browser to `localhost:8080` then displays the welcome page: ![Welcome pagel](/images/go-gae-bootstrap-quickstart/welcome.png)
