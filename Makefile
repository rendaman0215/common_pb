lint:
	@buf lint

format:
	@buf format -d

gen-api:
	@buf generate proto
