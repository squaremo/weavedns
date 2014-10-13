NAME:=weave-updater
IMAGE:=/tmp/$(NAME).tar
MAIN:=../

.DEFAULT: docker-image

all: docker-image

$(NAME): ../main.go
	go get -v -tags netgo $(MAIN) && go build -tags netgo -ldflags '-linkmode external -extldflags "-static"' -o $@ -v $(MAIN)

.PHONY: docker-image

docker-image: $(IMAGE)

$(IMAGE): Dockerfile $(NAME)
	sudo -E docker build -t zettio/$(NAME) .
	sudo -E docker save zettio/$(NAME) > $(IMAGE)

clean:
	-sudo docker rmi zettio/$(NAME)
	rm -f $(NAME) $(IMAGE)
