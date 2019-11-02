.PHONY: build, run, test, bench

build:
	@ cd app/backend && go build
	#@ docker build -t backend app/backend
	#@ docker run -p 80:8081 --rm backend

run: build
	@ cd app/backend && ./backend

test: build
	@ cd app/backend/prime && go test -v

bench:
	@ cd app/backend/prime && go test -bench=.
