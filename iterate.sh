#!/bin/bash
set -e
set -u
set -x
GOPATH=`pwd`:$GOPATH
PACKAGE_TOP=github.com/ncbray/lumen
echo "Generating code..."
go generate $PACKAGE_TOP/...
echo "Compiling tools..."
go install $PACKAGE_TOP/cmd/...
echo "Compiling DSL..."
bin/regenerate_lumen
go fmt $PACKAGE_TOP/...
