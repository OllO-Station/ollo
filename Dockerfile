FROM golang:1.19-alpine3.16 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

WORKDIR /ollo

# Add source files
COPY . .

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

RUN make build

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

COPY --from=builder /ollo/build/ /usr/local/bin/

# ENTRYPOINT ["ollod"]
