# -*- makefile -*-

new-entry:
	cd ./utils/new-entry/ && \
	go run .

build:
	cd ./utils/build-content-js/ && \
	go run . \
		--source=../../source \
		--content-html=../../content \
		--content-js=../../content.js
