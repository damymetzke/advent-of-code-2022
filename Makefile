build:
	for i in {1..15}; do go build -o bin/$$i days/$$i/main.go; done
