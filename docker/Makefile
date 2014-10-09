IMAGE:=/tmp/weavedns.tar

.DEFAULT: docker-image

all: docker-image

weavedns: ../server/*.go ../main.go
	$(eval TMP:=$(shell mktemp -d -t weave-build))
	cp -R ../server ../main.go ${TMP}/
	cp Dockerfile.build ${TMP}/Dockerfile
	docker build -t weavedns-build ${TMP}/
	docker run weavedns-build cat /home/weave/weavedns > $@
	chmod a+x $@

.PHONY: docker-image

docker-image: $(IMAGE)

$(IMAGE): Dockerfile weavedns
	sudo -E docker build -t zettio/weavedns .
	sudo -E docker save zettio/weavedns > $(IMAGE)

clean:
	-sudo docker rmi zettio/weavedns
	rm -f weavedns $(IMAGE)
