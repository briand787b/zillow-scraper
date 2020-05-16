run:
	docker-compose down --remove-orphans
	docker-compose build
	docker-compose config
	docker-compose up -d
	docker-compose logs -f
  
