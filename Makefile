GO ?= go
TPARSE ?= tparse

setup: install-pre-commit register-pre-commit install-tparse

test:
	$(GO) test -race -coverprofile=coverage.out -timeout 30s -json ./... | $(TPARSE) -all

install-tparse: # install tparse if not installed yet.
	@if ! command -v tparse > /dev/null; then \
		echo "Installing tparse"; \
		$(GO) install github.com/mfridman/tparse@latest; \
	fi

install-pre-commit:
	@if ! command -v pre-commit > /dev/null; then \
		echo "Installing pre-commit"; \
		pip3 install pre-commit; \
	fi

register-pre-commit:
	@pre-commit install --hook-type commit-msg