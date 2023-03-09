IMAGE := jackstock8904/cycling-blog
TAG := 1.2.6

.built-image: Dockerfile makefile
		docker build -t ${IMAGE} .
		touch .built-image

push: .built-image
		docker tag ${IMAGE} ${IMAGE}:${TAG}
		docker push ${IMAGE}:${TAG}
