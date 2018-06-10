#!/usr/bin/env bash
# Script for generating gocc packages.

set -eu

checkGocc() {
    # All files generated by gocc are kept under goccgen directory.
    src="./src"
    genDir="./goccgen"

    mkdir -p "$genDir"

    if [ -z "$GOPATH" ]; then
        echo "GOPATH environment variable is not set"
        exit 1
    else
        goccpath="$GOPATH"/bin/gocc
        if [ -f "$goccpath" ]; then
            "$goccpath" -a -zip -o "$genDir" "$src"/lang.bnf
        else
            echo "gocc is not properly installed"
            exit 1
        fi
    fi
}

checkGocc
