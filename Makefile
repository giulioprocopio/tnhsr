.PHONY: up down down-rm

up:
	python deploy_hook.py
	docker-compose up -d --build --force-recreate

down:
	docker-compose down

down-rm:
	docker-compose down -v --rmi all
	find . -name "*.volume" -type d -print0 | xargs -0 /bin/rm -rd