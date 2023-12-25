#!/bin/bash
CMD="./cmd"
BIN="${CMD}/$(basename ${PWD})"
MAIN="${CMD}/main.go"

case $1 in
    "linux-amd64")   export GOOS="linux"   GOARCH="amd64" ;;
    "darwin-amd64")  export GOOS="darwin"  GOARCH="amd64" ;;
    "freebsd-amd64") export GOOS="freebsd" GOARCH="amd64" ;;
    "windows-amd64") export GOOS="windows" GOARCH="amd64" ;;

    "clean")
        rm -vf "${BIN}"-*
        exit "$?" ;;

    *)
        echo -e "Supported platforms:"
        echo -e "linux-amd64\ndarwin-amd64\nfreebsd-amd64\nwindows-amd64"
        exit 0 ;;
esac


BIN="${BIN}-${GOOS}-${GOARCH}"

if [[ -e "${BIN}" ]]; then
    rm "${BIN}" && echo "Old build is deleted."
fi

/usr/bin/env go build -ldflags="-s -w" -o "${BIN}" "${MAIN}" && {
    echo -e "Build complete.\nUse ${BIN} to perform."
}

unset GOARCH GOOS
