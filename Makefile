.PHONY: golds mock

golds:
	golds -theme=dark -render-doclinks .

mock:
	mockery
