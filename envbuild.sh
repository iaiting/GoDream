#!/usr/bin/env bash

# ==============================================================================
#
# Author            : iaiting
#
# Contact           : iaiting@aliyun.com
#
# Generate Data     : 2018-02-17
#
# Copyright         : 本文件隶属iaiting，欢迎转载无须知会
#
# Desciption        : GoDream工程构建脚本，需要安装msys2，在msys2环境下运行该
#                     脚本
#
# ==============================================================================

GO=""
#===============================================================================
function check_env() {
    GO=`which go 2>/dev/null`
    if [ "X${GO}" == "X" ]; then
        echo "Please Config GO BIN Path..."
        return 1
    fi
}

#===============================================================================
function link_golang_org() {
    os_type=`${GO} env | grep GOOS | awk -F'[=]' '{print $2}'`

    if [ "${os_type}" = "windows" ]; then
        X_PATH=${GOPATH}\\src\\golang.org\\x
    else
        X_PATH=${GOPATH}/src/golang.org/x
    fi

    if [ ! -e ${X_PATH} ]; then
        mkdir -p ${X_PATH}
    fi

    if [ -e ${X_PATH}/tools ]; then
        cd ${X_PATH}/tools
        git pull
    else
        cd ${X_PATH}
        git clone https://github.com/golang/tools.git
    fi

    if [ -e ${X_PATH}/crypto ]; then
        cd ${X_PATH}/crypto
        git pull
    else
        cd ${X_PATH}
        git clone https://github.com/golang/crypto.git
    fi
}

#===============================================================================
function install_gotools() {
    ${GO} get -u -v -x github.com/uudashr/gopkgs/cmd/gopkgs

    # goreturns依赖tools，不能加-u，否则会去golang.org/x更新，而导致不成功
    ${GO} get -v -x sourcegraph.com/sqs/goreturns

    # 代码补全
    ${GO} get -u -v -x github.com/nsf/gocode

    # 代码跳转 未加-u选项
    ${GO} get -v -x github.com/rogpeppe/godef

    # 调试器
    ${GO} get -u -v -x github.com/derekparker/delve/cmd/dlv

    # 未加-u选项
    ${GO} get -v -x github.com/ramya-rao-a/go-outline
}

#===============================================================================
function main() {
    check_env
    if [ $? -ne 0 ]; then
        return 1
    fi

    link_golang_org
    if [ $? -ne 0 ]; then
        return 1
    fi

    install_gotools
    if [ $? -ne 0 ]; then
        return 1
    fi
}

#===============================================================================
main "$@"
