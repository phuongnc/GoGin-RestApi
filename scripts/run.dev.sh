#!/bin/sh

TAG="latest"
IMAGE="shapeservice"
echo "~~~~~ Starting build developement entity-server ${IMAGE} ${TAG} ~~~~~~~~~"
docker build -t $IMAGE:$TAG  -f ./build/.docker/dev.dockerfile .
echo '~~~~~ Finish build developement entity-server ~~~~~~~~~'
docker run -d -p 7171:7171 $IMAGE:$TAG
echo '~~~~~ Finish run developement entity-server ~~~~~~~~~'