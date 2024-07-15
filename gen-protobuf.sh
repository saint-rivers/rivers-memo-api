#!/bin/bash

protoc --experimental_allow_proto3_optional \
	--proto_path=$PWD/../rivers-memo-proto \
	-I=$PWD --go_out=$PWD \
	$PWD/../rivers-memo-proto/memo.proto
