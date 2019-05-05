# goCoupon [![Build Status][1]][2]

[1]: https://travis-ci.org/gabrielerzinger/goCoupon.svg?branch=master
[2]: https://travis-ci.org/gabrielerzinger/goCoupon

A Go API, backed by Redis, to manage coupons, promotional codes and referrals.

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Redis is, ofc a dep. You can easily get redis running using docker-compose:

```console
docker-compose up -d redis
```

Running it then should be as simple as:

```console
$ make build
$ ./bin/goCoupon
```

### Testing

Docker-compose is a testing dep. Running the tests should be simple as:
``make test``
