run:
	docker-compose down --remove-orphans
	docker-compose build
	docker-compose config
	docker-compose up -d
	docker-compose logs -f

update:
	docker-compose kill $(service)
	docker-compose build $(service)
	docker-compose up -d $(service)
	docker-compose logs -f
  
