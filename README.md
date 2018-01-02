# Gomematic: Web UI

[![Build Status](http://github.dronehippie.de/api/badges/gomematic/gomematic-ui/status.svg)](http://github.dronehippie.de/gomematic/gomematic-ui)
[![Stories in Ready](https://badge.waffle.io/gomematic/gomematic-api.svg?label=ready&title=Ready)](http://waffle.io/gomematic/gomematic-api)
[![Join the Matrix chat at https://matrix.to/#/#gomematic:matrix.org](https://img.shields.io/badge/matrix-%23gomematic%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gomematic:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/95976eab62b043c682ced6bda5b9021e)](https://www.codacy.com/app/gomematic/gomematic-ui?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomematic/gomematic-ui&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gomematic/gomematic-ui/server?status.svg)](http://godoc.org/github.com/gomematic/gomematic-ui/server)
[![Go Report](http://goreportcard.com/badge/github.com/gomematic/gomematic-ui)](http://goreportcard.com/report/github.com/gomematic/gomematic-ui)
[![](https://images.microbadger.com/badges/image/gomematic/gomematic-ui.svg)](http://microbadger.com/images/gomematic/gomematic-ui "Get your own image badge on microbadger.com")


**This project is under heavy development, it's not in a working state yet!**

TBD


## Build

This project requires NodeJS to build the sources, the installation of NodeJS won't be covered by those instructions. To build the sources just execute the following command after NodeJS setup:

```
yarn install
yarn run build
```

If you also want to publish it as a single binary with our server based on Go make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring you have to use a Go version `>= 1.6`

```bash
go get -d github.com/gomematic/gomematic-ui
cd $GOPATH/src/github.com/gomematic/gomematic-ui
make generate build

./gomematic-ui -h
```

With the `make generate` command we are embedding all the static assets into the binary so there is no need for any webserver or anything else beside launching this binary.


## Development

To start developing on this UI you have to execute only a few commands. To setup a NodeJS environment or even a Go environment is out of the scope of this document. To start development just execute those commands:

```
yarn install
yarn run watch

./gomematic-ui server --static dist/static/
```

The development server reloads the used assets on every request. So in order to properly work with it you need to start the API separately. After launching this command on a terminal you can access the web interface at [http://localhost:9000](http://localhost:9000)


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
