docker: 
	docker-compose -f docker-compose.dev.yml up

down:
	docker-compose down
	docker ps -a

all: 
	docker-compose -f docker-compose.yml up

alldown:
	docker-compose down
	docker ps -a

.PHONY: docker down all alldown