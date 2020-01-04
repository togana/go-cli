.PHONY: build
build:
		rm -rf ./bin && \
		docker-compose up --build && \
		docker cp `docker-compose ps -q cli`:/app/darwin_amd64/ ./bin
