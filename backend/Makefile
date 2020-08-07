NAME					=	cbm-api

API_DIR					=	api

.DEFAULT: api-start

######################################################################################
# API
api-start:
	@docker-compose up -d --build

api-run:
	@docker-compose up --build

api-dev:
	watchexec --restart --filter "$(pwd)/*" --watch . "docker-compose -f docker-compose.dev.yml up --build"

api-test:
	@docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

api-unit-test: api-start
	-@go test ./...
	@make --no-print-directory api-stop

api-stop:
	@docker-compose down --remove-orphans

api-clean-stop:
	@docker-compose down --remove-orphans --volumes

api-clean-run: api-clean-stop api-run

api-clean-start: api-clean-stop api-start


.PHONY: api-main api-run api-test api-stop api-clean-stop



######################################################################################
# Misc
clean: api-clean-stop
	@echo "y" | docker system prune

fclean: clean
	@echo "y" | docker system prune -a --volumes --remove-orphans

restart: api-stop
restart: clean
restart: .DEFAULT



.PHONY: clean fclean restart
