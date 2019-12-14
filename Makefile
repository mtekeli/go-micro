ROOT_DIR:=${shell dirname ${realpath ${lastword ${MAKEFILE_LIST}}}}

.PHONY: build, run, deploy, test, bench

build:
	@ cd ${ROOT_DIR}/app/backend && docker-compose build
	@ cd ${ROOT_DIR}/app/frontend && docker-compose build
	
run: build
	@ cd ${ROOT_DIR}/app/backend && docker-compose up &
	@ cd ${ROOT_DIR}/app/frontend && docker-compose up &

deploy: build
	#@ docker network create --scope=swarm --driver=bridge --subnet=172.22.0.0/16 --gateway=172.22.0.1 backend_bridge
	@ docker stack deploy --compose-file app/backend/docker-compose.yml dev_backend
	@ docker stack deploy --compose-file app/frontend/docker-compose.yml dev_frontend

stop:
	#@ cd ${ROOT_DIR}/app/frontend && docker-compose down
	#@ cd ${ROOT_DIR}/app/backend && docker-compose down
	@ docker stack rm dev
	@ docker stack rm dev_frontend
	@ docker stack rm dev_backend

test:
	@ cd ${ROOT_DIR}/app/backend/prime && go test -v

bench:
	@ cd ${ROOT_DIR}/app/backend/prime && go test -bench=.
