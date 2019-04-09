# Build Stage
FROM lacion/alpine-golang-buildimage:1.10.3 AS build-stage

LABEL app="build-goCoupon"
LABEL REPO="https://github.com/gabrielerzinger/goCoupon"

ENV PROJPATH=/go/src/github.com/gabrielerzinger/goCoupon

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/gabrielerzinger/goCoupon
WORKDIR /go/src/github.com/gabrielerzinger/goCoupon

RUN make build-alpine

# Final Stage
FROM lacion/apine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/gabrielerzinger/goCoupon"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/goCoupon/bin

WORKDIR /opt/goCoupon/bin

COPY --from=build-stage /go/src/github.com/gabrielerzinger/goCoupon/bin/goCoupon /opt/goCoupon/bin/
RUN chmod +x /opt/goCoupon/bin/goCoupon

# Create appuser
RUN adduser -D -g '' goCoupon
USER goCoupon

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/goCoupon/bin/goCoupon"]
