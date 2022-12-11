build:
	for i in {1..11}; do go build -o bin/$$i days/$$i/main.go; done
