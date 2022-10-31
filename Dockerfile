ARG GO_VERSION="1.18"
ARG RUNNER_IMAGE="gcr.io/distroless/static"

# Builder
FROM golang:${GO_VERSION}-alpine as builder

ARG GIT_VERSION
ARG GIT_COMMIT

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

WORKDIR /ollo
COPY go.mod go.sum ./
RUN target=/root/.cache/go-build \
    target=/root/go/pkg/mod \
    go mod download

COPY . .

RUN target=/root/.cache/go-build \
    target=/root/go/pkg/mod \
    go build \
      -mod=readonly \
      -tags "netgo,ledger,muslc" \
      -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name="ollo" \
              -X github.com/cosmos/cosmos-sdk/version.AppName="ollod" \
              -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
              -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
              -X github.com/cosmos/cosmos-sdk/version.BuildTags='netgo,ledger,muslc' \
              -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
      -trimpath \
      -o /ollo/build/ollod \
      /ollo/cmd/ollod/main.go

# Runner
FROM ${RUNNER_IMAGE}

COPY --from=builder /ollo/build/ollod /bin/ollod

ENV HOME /ollo
WORKDIR $HOME

EXPOSE 26656
EXPOSE 9091
EXPOSE 26657
EXPOSE 1317

ENTRYPOINT ["ollod"]
