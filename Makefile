# .PHONY: all
.DEFAULT_GOAL := generate_all

generate_all: clone generate clean

clone:
	git clone https://github.com/binary-com/deriv-developers-portal.git

generate: 
	go run bin/generate.go

clean:
	rm -rf deriv-developers-portal