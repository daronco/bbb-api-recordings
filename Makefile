PWD=$(shell pwd)

image-dev:
	docker build -t bigbluebutton/bbb-api-recordings:dev .

run-dev:
	docker run --rm -ti -v ${PWD}:/go/src/github.com/bigbluebutton/bbb-api-recordings \
		-p 8081:8081 --name bbb-api-recordings bigbluebutton/bbb-api-recordings:dev
