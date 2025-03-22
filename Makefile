docker-build:
	docker build --tag localhost:3000/shooters/user-api:latest .

docker-push:
	docker push localhost:3000/shooters/user-api:latest

docker-deploy:
	@make docker-build
	@make docker-push