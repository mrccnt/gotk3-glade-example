#!/usr/bin/env bash

script=$(readlink -f "$0")
scriptpath=$(dirname "${script}")

GOBIN="${scriptpath}/bin" go install -v -tags gtk_3_18 .
