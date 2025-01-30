DATABASE_PWD:=test@123
DATABASE_URL:=postgres://postgres:${DATABASE_PWD}@localhost:5432/postgres

# This docker command build and run a docker container
# that connects to the host postgres database instance.
.PHONY: docker-build-run
docker-build-run:
	docker build -t api-forex:v0.0.1 .
	docker run \
		--network=host \
		-e DATABASE_URL=${DATABASE_URL} \
		-p 8080:8080 \
		--restart unless-stopped \
		api-forex:v0.0.1