ARG GOVERSION
FROM golang:${GOVERSION:-1.22.1} AS build_ksmapi

COPY ./ /project
WORKDIR /project

ARG APPVERSION="0.0.0-DEV"

RUN make build TAG=${APPVERSION} \
    && chmod +x /project/dist/ksmapi

# build the production final image
FROM scratch as PROD

ARG APPVERSION="0.0.0-DEV"

LABEL org.opencontainers.image.maintainer="Denis Rendler <connect@rendler.net>" \
    org.opencontainers.image.version=$APPVERSION \
    org.opencontainers.image.description="REST API app for Keeper Secrets Manager" \
    org.opencontainers.image.source="https://github.com/Hexagonal-Software/ksm-api" \
    org.opencontainers.image.url="https://github.com/Hexagonal-Software/ksm-api" \
    org.opencontainers.image.authors="Denis Rendler <connect@rendler.net>"

COPY --from=build_ksmapi /project/dist/ksmapi /ksmapi/ksmapi
COPY --from=build_ksmapi /project/dist/ksmapi.yaml /ksmapi/ksmapi.yaml
COPY --from=build_ksmapi /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /ksmapi

ARG http_port=8086
EXPOSE ${http_port}/tcp

ENTRYPOINT ["/ksmapi/ksmapi"]
CMD ["server"]
