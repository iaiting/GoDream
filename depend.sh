#!/usr/bin/env bash

################################################################################
function main() {
    go get -u -v -x github.com/gin-gonic/gin
    go get -v -x github.com/tjfoc/gmsm/sm2
}

################################################################################
main "$@"
