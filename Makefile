setup: install-pre-commit register-pre-commit

install-pre-commit:
	@if ! command -v pre-commit > /dev/null; then \
		echo "Installing pre-commit"; \
		pip3 install pre-commit; \
	fi

register-pre-commit:
	@pre-commit install --hook-type commit-msg