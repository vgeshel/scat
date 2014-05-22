#!/bin/sh

set -e

go build

docker build -t vgeshel/scat .


