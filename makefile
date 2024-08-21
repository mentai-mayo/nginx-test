
.PHONY: all build run clean

all: clean build run

build:
	docker build -t nginx-test . --progress plain

run:
	docker run -it -p 8808:80 --name nginx-test nginx-test

clean:
	-docker stop nginx-test
	-docker rm nginx-test
	-docker rmi nginx-test
