.PHONY: up down down-rm

up:
	docker-compose up -d --build --force-recreate

down:
	docker-compose down

down-rm:
	docker-compose down -v --rmi all