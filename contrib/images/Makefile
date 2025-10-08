all: postworldd-env

postworldd-env: postworldd-rmi
	docker build --tag cosmos/postworldd -f postworldd-env/Dockerfile \
		$(shell git rev-parse --show-toplevel)

postworldd-rmi:
	docker rmi cosmos/postworldd 2>/dev/null; true

.PHONY: all postworldd-env postworldd-rmi
