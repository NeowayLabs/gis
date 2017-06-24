.PHONY: test

all: test

test: deps
	cd _testdata/geodata && git pull origin master
	go test -v ./...

deps: ./_testdata ./_testdata/geodata

./_testdata:
	mkdir -p _testdata

./_testdata/geodata:
	cd ./_testdata && git clone https://github.com/NeowayLabs/geodata.git
