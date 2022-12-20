build:
	for i in {1..16}; do pushd days/$$i >/dev/null; go build -o ../../bin/$$i main.go; popd >/dev/null; done
