include flip.env
export $(shell sed 's/=.*//' flip.env)

	#docker-compose run --rm --entrypoint="./appmain withdraw --account_number=444444 --amount=5555 --remark=hii --bank_code=bca" app

build:
	docker build . -t my_market
start:
	docker-compose up
stop:
	docker-compose stop
clean:
	docker-compose down
withdraw:
	c=$$(echo "docker-compose run --rm --entrypoint=\"./appmain withdraw --account_number=$$ACCOUNT_NUMBER --amount=$$AMOUNT --remark=$$REMARK --bank_code=$$BANK_CODE\" app") && echo $$c && bash -c "$$c"
heroku-config:
	heroku config:set \
	DBHost=$$(echo $$DBHost) \
	DBPort=$$(echo $$DBPort) \
	DBUser=$$(echo $$DBUser) \
	DBPassword=$$(echo $$DBPassword) \
	DBName=$$(echo $$DBName) \
	DBSsl=$$(echo $$DBSsl) \
	FlipHost=$$(echo $$FlipHost) \
	FlipSecret=$$(echo $$FlipSecret) \
	ListenPort=$$(echo $$ListenPort)
heroku-create-db:
	heroku addons:create heroku-postgresql:hobby-dev
	heroku config:get DATABASE_URL
heroku-deploy: heroku-config
	heroku container:login
	heroku container:push web
	heroku container:release web