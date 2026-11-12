.PHONY: hooks

# Active les hooks versionnés dans .githooks/ (à lancer une fois après le clone)
hooks:
	git config core.hooksPath .githooks
	chmod +x .githooks/pre-commit .githooks/pre-push
