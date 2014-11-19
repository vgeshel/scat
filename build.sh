#!/bin/bash
im=vgeshel/scat
set -e

git submodule update --init
# build the executable
docker build -t $im-stage .
# build the small image using the executable and other files from the staging image
bash -evx build-small-image.sh $im $im-stage /usr/src/app/bin/scat
# stage image no longer needed
docker rmi $im-stage
