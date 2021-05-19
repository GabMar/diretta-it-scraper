FROM golang:1.16

RUN go get -v \
        golang.org/x/tools/gopls \
        honnef.co/go/tools/... \
        golang.org/x/lint/golint \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/go-delve/delve/cmd/dlv \
    && go clean -modcache


COPY --from=golangci/golangci-lint:v1.31-alpine /usr/bin/golangci-lint /usr/bin/golangci-lint
