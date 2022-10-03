.PHONY: run
run:
	./scrypts/run_app.sh

.PHONY: build
build:
	go build -o build/rona cmd/http/main.go 

.PHONY: build-and-run
build-and-run: build
	./build/rona

.PHONY: remove-build
remove-build:
	rm -r build


.PHONY: run-mongo
run-mongo:
	./scrypts/run_mongo.sh
