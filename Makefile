build:
	for i in {1..7}; do go build -o bin/$$i days/$$i/main.go; done
