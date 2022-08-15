GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test
SWAG=swag
CODE_PATH=./src
PREFIX=vigilant_waddle_

usage:
	@printf "\nMake usage:\n"
	@printf "\tmake run \t- run the server\n"
	@printf "\tmake docs \t- generate swagger documentation\n"
	@printf "\tmake rebuild_db - rebuild the database\n"
	@printf "\tmake stop \t- stop the server\n"
	@printf "\tmake mock \t- generate application mocks\n"
	@printf "\tmake test \t- run tests\n"
	@printf "\tmake e2e-test \t- run e2e tests\n"

.PHONY: docs
docs:
	$(SWAG) init

.PHONY: run
run:
	$(SWAG) init
	docker-compose up -d $(PREFIX)api_db
	docker-compose up -d $(PREFIX)api_client
	docker-compose up -d $(PREFIX)api_docs

.PHONY: rebuild_db
rebuild_db:
	@./local-scripts/destroy-db-container;\
	docker-compose up -d $(PREFIX)api_db

.PHONE: stop
stop:
	docker-compose down

.PHONY: mock
mock:
	./local-scripts/install-mockery
	source ~/.bashrc
	mockery --all

.PHONY: test
test:
	cd $(CODE_PATH) && mkdir -p ./tests
	cd $(CODE_PATH) && ENVIRONMENT=testing $(GOTEST) -coverprofile=./tests/coverage-report.out ./...
	cd $(CODE_PATH) && $(GOCOVER) -func=./tests/coverage-report.out
	cd $(CODE_PATH) && $(GOCOVER) -html=./tests/coverage-report.out

.PHONY: e2e-test
e2e-test:
	@./local-scripts/integration-tests;\
	exit $$?;
