POSTGRES := postgres
POSTGRES_IMAGE := postgres:9.4
POSTGRES_PASSWORD := password

run-pg:
	@docker run \
	  -d \
	  -p 5432:5432 \
	  --name $(POSTGRES) \
	  $(POSTGRES_IMAGE)

stop-pg:
	@docker stop $(POSTGRES)
	@docker rm $(POSTGRES)

.PHONY: run-pg stop-pg
