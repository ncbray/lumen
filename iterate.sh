#!/bin/bash
set -e
set -u
GOPATH=`pwd`:$GOPATH
PACKAGE_TOP=github.com/ncbray/reboot
echo "Generating code..."
go generate $PACKAGE_TOP/...
echo "Compiling tools..."
go install $PACKAGE_TOP/cmd/...
echo "Testing"
bin/playground src/$PACKAGE_TOP/ast/ast.dsl
