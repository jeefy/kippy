TOPDIR=$(PWD)
WHOAMI=$(shell whoami)
SENDGRID_API_KEY=$(shell cat jeefy/sendgrid.token)
STRIPE_API_KEY=$(shell cat jeefy/stripe.token.dev)

build:
	go build -o bin/kippy cmd/main.go

run:
	go run cmd/main.go --discordWebhook="${DISCORD_WEBHOOK_URL}"

image:
	docker build -t ${WHOAMI}/kippy -f Dockerfile .

image-push: image
	docker push ${WHOAMI}/kippy