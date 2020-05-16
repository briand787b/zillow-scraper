run:
	docker-compose down
	docker-compose build
	docker-compose config
	docker-compose up -d
	docker-compose logs -f
  
