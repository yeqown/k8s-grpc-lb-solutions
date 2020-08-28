#!/bin/bash

images="
gcr.io/linkerd-io/controller:stable-2.7.1
gcr.io/linkerd-io/proxy:stable-2.7.1
gcr.io/linkerd-io/web:stable-2.7.1
gcr.io/linkerd-io/grafana:stable-2.7.1
gcr.io/linkerd-io/proxy-init:v1.3.2
"

for image in $images
    do
        aliyunImage=registry.cn-hangzhou.aliyuncs.com/doji-linkerd-io/$(echo $image | cut -d '/' -f 3)
        echo $image, $aliyunImage
        docker pull $aliyunImage
        docker tag $aliyunImage $image
    done