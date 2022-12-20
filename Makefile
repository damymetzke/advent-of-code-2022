build:
	for i in {1..17}; do pushd days/$$i >/dev/null; go build -o ../../bin/$$i main.go; popd >/dev/null; done
