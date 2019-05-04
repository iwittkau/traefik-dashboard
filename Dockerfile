FROM alpine:3.9

COPY traefik-dashboard /
COPY templates /templates

EXPOSE 80

ENV GIN_MODE=release

ENTRYPOINT [ "/traefik-dashboard" ]