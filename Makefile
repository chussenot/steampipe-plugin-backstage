.PHONY: help install test docs
# This Makefile is designed to automate the build process for a project.
# It includes targets for compiling the code, cleaning up build artifacts,
# and other common tasks. The default target is set to 'help', which provides
# a summary of available targets and their descriptions.
#
# Targets:
# - install: Installs the built executables and other necessary files to the system.
# - test: Runs tests for the project.
# - docs: Generates documentation for the project.

install: ## Install the plugin
	rm -rf ~/.steampipe/plugins/local/steampipe-plugin-backstage
	go build -o ~/.steampipe/plugins/local/steampipe-plugin-backstage/steampipe-plugin-backstage.plugin *.go

.PHONY: build
build: ## Build the plugin
	go build -o ./build/steampipe-plugin-backstage.plugin *.go

test:
	go test -v ./...

docs:
	go generate ./...

config: ## Copy the backstage config to the steampipe config folder
	cp config/backstage.spc $(HOME)/.steampipe/config/backstage.spc

.DEFAULT_GOAL := help

help:
	@echo "Available targets:"
	@echo "  install: Install the plugin"
	@echo "  build: Build the plugin"
	@echo "  test: Run tests"
	@echo "  docs: Generate documentation"
	@echo "  config: Copy the backstage config to the steampipe config folder"
