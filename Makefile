.PHONY: snapshot
snapshot:
	goreleaser release --snapshot --rm-dist

.PHONY: run-docker
run-docker:
	docker run \
		-d \
		--env GIN_MODE=debug \
		--rm \
		--name traefik-dashboard \
		--network traefik_proxy \
		-l "traefik.frontend.rule=Host:dashboard.docker.localhost" \
		iwittkau/traefik-dashboard:latest \
		-h traefik:8080

.PHONY: run
run:
	go generate ./...
	go run cmd/traefik-dashboard/main.go -t "templates" -a ":8080"