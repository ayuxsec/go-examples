default: make-push

.PHONY: make-push
make-push:
	git add -A
	git config user.name "ayuxsec" && git config user.email "ayuxsec@proton.me"
	git commit -m "make-push"
	git push
