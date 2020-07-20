install:
	docker-compose build chrome-app
	docker-compose run --rm --no-deps chrome-app npm install

run:
	docker-compose down --remove-orphans
	docker-compose build
	docker-compose config
	docker-compose up -d
	docker-compose logs -f

# opens up terminal - deps must be added manually
add-go-deps:
	docker image build \
		-f ./backend/go/api/Dockerfile \
		--target builder \
		-t go-deps \
		./backend/go
	docker container run \
		-it \
		--rm \
		-v ${PWD}/backend/go:/go/app/ \
		go-deps /bin/bash


update:
	docker-compose kill $(service)
	docker-compose build $(service)
	docker-compose up -d $(service)
	docker-compose logs -f
  
test:
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		config
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down \
			--remove-orphans
	-docker volume rm redis_test_data
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run --rm backend-go-test
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down \
			--remove-orphans

test-interactive:
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		config
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down \
			--remove-orphans
	-docker volume rm redis_test_data
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run \
		-v ${PWD}/backend/go:/go/app \
		backend-go-test /bin/bash