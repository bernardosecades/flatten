.PHONY: help dependencies up start stop restart status ps clean and execute tests

ps:
	docker-compose ps
up:
	docker-compose up --build
down:
	docker-compose down
test:
	docker-compose exec web bash -c "go clean -testcache  && go test ./... -tags=unit"
