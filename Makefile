all: counter pounder

counter:
	cd counter && \
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/counter.${GOARCH} counter.go && \
	docker image build --build-arg DFROM=${DFROM} --build-arg ARCH=${GOARCH} --tag staranto/simple-http-counter:${GOARCH} . && \
	docker image push staranto/simple-http-counter:${GOARCH}; \
	cd ..

pounder:
	cd pounder && \
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/pounder.${GOARCH} pounder.go && \
	docker image build --build-arg DFROM=${DFROM} --build-arg ARCH=${GOARCH} --tag staranto/simple-http-pounder:${GOARCH} . && \
	docker image push staranto/simple-http-pounder:${GOARCH}; \
	cd ..