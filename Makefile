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