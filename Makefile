APP?=qsm
RELEASE?=$(shell python version.py get)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell powershell get-date -format "{yyyy-mm-dd_HH:mm:ss}")
PROJECT?=github.com/Jarover/qsm

clean:
	rm -f ${APP}
	rm -f ${APP}.exe


buildwin: clean
	python version.py inc-patch
	GOOS=windows go build \
				-o ${APP}.exe \
                -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
                -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
                cmd/${APP}/main.go


buildlinux:	
	rm -f ${APP} 
	GOOS=linux go build \
				-o ${APP} \
                -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
                -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
				$(go list -m)/cmd/${APP}

deploy: build
	./deploy.sh

run:	build
	./${APP} -f dev.json

test:
	go test -v -race ./...