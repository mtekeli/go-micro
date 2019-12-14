.PHONY: build, run, deploy, test, bench

build:
	@ docker-compose build
	
run: build
	@ docker-compose up &

deploy: build
	@ docker stack deploy --compose-file docker-compose.yml dev

stop:
	@ docker-compose down
	@ docker stack rm dev

test:
	@ cd prime && go test -v

bench:
	@ cd prime && go test -bench=.
