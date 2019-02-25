#!/bin/sh

cd $(dirname $BASH_SOURCE)

mkdir -p build/ 
go build -o build/crawl