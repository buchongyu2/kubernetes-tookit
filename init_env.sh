#!/bin/bash

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go env -w CGO_ENABLED=0
go mod download
