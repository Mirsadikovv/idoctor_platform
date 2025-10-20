run:
	go run -tags dev .

docker-run:
	docker compose --env-file .env.dev -f docker-compose-dev.yml up --build -d

docker-down:
	docker compose -f docker-compose-dev.yml down -v 

swag_init:
	@echo "Generating Swagger documentation..."
	@find src/module -type f -name "cmd.go" | while read -r file; do \
		dir=$$(dirname "$$file"); \
		module_name=$$(basename "$$dir"); \
		echo "Processing $$module_name..."; \
		(cd "$$dir" && swag init -pd --ot json --o "../../docs/$$module_name" -g cmd.go > /dev/null 2>&1) || echo "Error in $$module_name"; \
	done
	@echo "Swagger generation complete."

swag_init_werr:
	@echo "Generating Swagger documentation..."
	@error_count=0; \
	find src/module -type f -name "cmd.go" | while read -r file; do \
		dir=$$(dirname "$$file"); \
		module_name=$$(basename "$$dir"); \
		echo "Processing $$module_name..."; \
		error_output=$$(cd "$$dir" && swag init -pd --ot json --o "../../docs/$$module_name" -g cmd.go 2>&1); \
		if [ $$? -ne 0 ]; then \
			echo "Error in $$module_name:"; \
			echo "$$error_output"; \
			echo "File: $$file"; \
			echo "----------------------------------------"; \
			error_count=$$((error_count + 1)); \
		else \
			echo "Successfully processed $$module_name"; \
		fi; \
	done; \
	if [ $$error_count -eq 0 ]; then \
		echo "Swagger generation complete. No errors detected."; \
	else \
		echo "Swagger generation completed with $$error_count error(s)."; \
	fi

pull-subtree:
	git subtree pull --prefix=dmi_bot https://git.sriss.uz/mehnat/dmi_bot.git dev --squash
push-subtree:
	git subtree push --prefix=dmi_bot https://git.sriss.uz/mehnat/dmi_bot.git dev