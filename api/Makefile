NAME					=	cbm-api

API_DIR					=	api

.DEFAULT: api-start


######################################################################################
# API
api-start:
	@docker-compose up -d --build

api-run:
	@docker-compose up --build

api-test:
	@docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

api-unit-test: api-start
	-@go test ./...
	@make --no-print-directory api-stop

api-stop:
	@docker-compose down --remove-orphans --volumes


.PHONY: api-main api-run api-test api-stop



######################################################################################
# Misc
clean: api-stop
	@echo "y" | docker system prune

fclean: clean
	@echo "y" | docker system prune -a --volumes

restart: api-stop
restart: clean
restart: .DEFAULT



.PHONY: clean fclean restart
