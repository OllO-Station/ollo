#!/usr/bin/env bash

set -eo pipefail

out_dir="protogo"
# move the vendor folder to a temp dir so that go list works properly
temp_dir="f29ea6aa861dc4b083e8e48f67cce"
if [ -d vendor ]; then
  mv ./vendor ./$temp_dir
fi

# Get the path of the cosmos-sdk repo from go/pkg/mod
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)

# move the vendor folder back to ./vendor
if [ -d $temp_dir ]; then
  mv ./$temp_dir ./vendor
fi

proto_dirs=$(find . \( -path ./third_party -o -path ./vendor \) -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

rm -rf github.com; rm -rf cosmos; rm -rf ollo
rm -rf ${out_dir}
rm -rf cosmossdk.io

mkdir ${out_dir}
for dir in $proto_dirs; do
  # generate protobuf bind
  protoc \
  -I "proto" \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
  $(find "${dir}"  -name '*.proto')

  # generate grpc gateway
  protoc \
  -I "proto" \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  -I "$cosmos_sdk_dir/proto" \
  --grpc-gateway_out=logtostderr=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')
done
mv ollo ${out_dir}/ollo


for dir in $(find third_party/proto/cosmos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  for file in $(find "${dir}" -maxdepth 1  -name '*.proto'); do
  echo "Generating $file"
  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --gocosmos_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. $file 

  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --grpc-gateway_out=logtostderr=true:. $file
  done
done
rm -rf cosmossdk.io
rm -rf cosmos
mv github.com/cosmos/cosmos-sdk ${out_dir}/cosmos

for dir in $(find third_party/proto/cosmwasm -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  for file in $(find "${dir}" -maxdepth 1  -name '*.proto'); do
  echo "Generating $file"
  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --gocosmos_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. $file 

  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --grpc-gateway_out=logtostderr=true:. $file
  done
done
mv github.com/CosmWasm/wasmd ${out_dir}/cosmwasm

for dir in $(find third_party/proto/ibc -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  for file in $(find "${dir}" -maxdepth 1  -name '*.proto'); do
  echo "Generating $file"
  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --gocosmos_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. $file 

  protoc \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  --grpc-gateway_out=logtostderr=true:. $file
  done
done
mv github.com/cosmos/ibc-go ${out_dir}/ibc

rm -rf github.com

