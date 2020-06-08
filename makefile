IMAGE := jackstockley89/golangwebpage
TAG := 1.1

.built-image: Dockerfile makefile
		docker build -t ${IMAGE} .
		touch .built-image

push: .built-image
		docker tag ${IMAGE} ${IMAGE}:${TAG}
		docker push ${IMAGE}:${TAG}
