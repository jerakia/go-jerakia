PWD := $(shell pwd)
JERAKIA_ETC_DIR="$(PWD)/acceptance/fixtures/etc/jerakia"
JERAKIA_VAR_DIR="$(PWD)/acceptance/fixtures/var/lib/jerakia"
HIERA_VAR_DIR="$(PWD)/acceptance/fixtures/var/lib/hiera"
#JERAKIA_IMAGE="crayfishx/jerakia:latest"
JERAKIA_IMAGE="jerakia/jerakia:latest"

test:
	go test -v ./testing

.ONESHELL:
docker:
	for i in $$(docker ps -q -f name=jerakia-server); do docker rm -f $$i; done
	docker pull jerakia/jerakia:$(JERAKIA_VERSION)
	docker run -p 9843:9843 -d --rm --name jerakia-server \
	           -v $(JERAKIA_ETC_DIR):/etc/jerakia \
	           -v $(JERAKIA_VAR_DIR):/var/lib/jerakia \
	           -v $(HIERA_VAR_DIR):/var/lib/hiera \
	           $(JERAKIA_IMAGE)

.ONESHELL:
testacc: docker
	sleep 5
	export JERAKIA_TOKEN=$$(docker exec jerakia-server jerakia token create myapp --quiet)
	export JERAKIA_URL="http://localhost:9843/v1"
	JERAKIA_ACC=1 go test -v ./acceptance -run="$(TEST)"
