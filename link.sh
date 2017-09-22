#!/usr/bin/env bash

for dir in `find vendor -mindepth 1 -maxdepth 1 -type d`
do
    segment=$(echo ${dir} | awk -F/ '{print $NF}')
    [ -d src/${segment} ] && rm -r src/${segment}
    ln -sr vendor/${segment} src/${segment}
done
