FROM golang:1.19-alpine as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

WORKDIR /ollo

# Add source files
COPY . .

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOFLAGS="-buildvcs=false"

RUN go mod tidy
RUN make install

# ----------------------------

FROM alpine:3.16

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660
# api port
EXPOSE 1317

EXPOSE 9090

EXPOSE 6060

COPY --from=builder /ollo/ollod /usr/local/bin/

# ENTRYPOINT ["ollod"]
