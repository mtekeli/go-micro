.PHONY: build, run, test, bench

build:
	@ docker-compose build
	
run: build
	@ docker-compose up &

stop:
	docker-compose down

test:
	@ cd app/backend/prime && go test -v

bench:
	@ cd app/backend/prime && go test -bench=.
