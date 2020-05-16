test:
	GO111MODULE=on go test ./...

build: test
	GO111MODULE=on go build

run: build
	./util

# can be more of course (see runp)
#PLATFORMS := linux/arm linux/amd64
PLATFORMS := linux/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: test $(PLATFORMS)

$(PLATFORMS):
	# Build multiplatform images
	docker build --build-arg GOOS=$(os) --build-arg GOARCH=$(arch) -t util-$(os)-$(arch) .

	# Push image to public registry - hub.docker.com
	docker login
	docker tag util-$(os)-$(arch):latest reisinge/util-$(os)-$(arch):latest
	docker push reisinge/util-$(os)-$(arch):latest
