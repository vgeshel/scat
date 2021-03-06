#!/bin/bash

image_name=$1
source_image=$2
executable=$3

scratch=scratch-$$

set -evx

docker run -it --name $scratch --entrypoint /bin/true $source_image

dir=$(mktemp -d)

pushd $dir

docker cp $scratch:/etc/ssl/certs ./etc/ssl/
docker cp $scratch:$executable .`dirname $executable`

docker rm $scratch

cat > Dockerfile <<EOF
FROM scratch
ADD . /
ENTRYPOINT ["$executable"]
EOF

docker build -t $image_name .

popd

rm -rf $dir
