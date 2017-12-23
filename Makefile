cpm: src/*
	go build -o cpm ./src/

deps:
	go get github.com/urfave/cli
