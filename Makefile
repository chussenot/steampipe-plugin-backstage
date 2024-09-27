.PHONY: help install
# This Makefile is designed to automate the build process for a project.
# It includes targets for compiling the code, cleaning up build artifacts,
# and other common tasks. The default target is set to 'help', which provides
# a summary of available targets and their descriptions.
#
# Targets:
# - install: Installs the built executables and other necessary files to the system.

install: ## Install the plugin
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/chussenot/backstage@latest/steampipe-plugin-backstage.plugin *.go
