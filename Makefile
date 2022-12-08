build:
	for i in {1..8}; do go build -o bin/$$i days/$$i/main.go; done
