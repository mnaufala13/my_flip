include flip.env
export $(shell sed 's/=.*//' flip.env)

build:
	docker build . -t my_market
start:
	docker-compose up
stop:
	docker-compose stop
clean:
	docker-compose down