NAME					=	cbm-api

API_DIR					=	api

.DEFAULT: api-start


######################################################################################
# API
api-start:
	@cd $(API_DIR) && docker-compose up -d --build

api-run:
	@cd $(API_DIR) && docker-compose up --build

api-test:
	@cd $(API_DIR) && docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

api-unit-test: api-start
	-@cd $(API_DIR) && go test ./...
	@make --no-print-directory api-stop

api-stop:
	@cd $(API_DIR) && docker-compose down --remove-orphans --volumes


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
