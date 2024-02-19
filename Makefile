COVER_FILEANME=cover.out
ALL_PKGS=./...

test:
	go test $(ALL_PKGS)

generate-coverage:
	go test -coverprofile $(COVER_FILEANME) $(ALL_PKGS) 
	
coverage-web: generate-coverage
	go tool cover -html=$(COVER_FILEANME)

coverage-terminal: generate-coverage
	go tool cover -func=$(COVER_FILEANME)

 