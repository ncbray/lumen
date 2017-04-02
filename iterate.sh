#!/bin/bash
set -e
set -u
GOPATH=`pwd`:$GOPATH
PACKAGE_TOP=github.com/ncbray/reboot
echo "Generating code..."
go generate $PACKAGE_TOP/...
echo "Compiling tools..."
go install $PACKAGE_TOP/cmd/...
echo "Compiling DSL..."
bin/regenerate
go fmt $PACKAGE_TOP/...
