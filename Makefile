install:
	docker-compose up -d
	docker exec -it application bash

stop:
	docker-compose down