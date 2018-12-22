all: counter pounder

counter:
	cd counter && \
	GOOS=linux GOARCH=${GOARCH} go build -o build/counter.${GOARCH} counter.go && \
	docker image build --build-arg ARCH=${GOARCH} --tag staranto/simple-http-counter:${GOARCH} . && \
	echo docker image push staranto/simple-http-counter:${GOARCH}; \
	cd ..

pounder:
	cd pounder && \
	GOOS=linux GOARCH=${GOARCH} go build -o build/pounder.${GOARCH} pounder.go && \
	docker image build --build-arg ARCH=${GOARCH} --tag staranto/simple-http-pounder:${GOARCH} . && \
	echo docker image push staranto/simple-http-pounder:${GOARCH}; \
	cd ..