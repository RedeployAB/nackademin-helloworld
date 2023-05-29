# golang:1.20.4-alpine3.17 SHA1 digest.
FROM golang@sha256:913de96707b0460bcfdfe422796bb6e559fc300f6c53286777805a9a3010a5ea as builder

ARG BIN=endpoint

ENV USER=${BIN}
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nohome" \
    --no-create-home \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"


WORKDIR /src/${BIN}
COPY . .

RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build \
    -o build/${BIN} \
    -ldflags="-s -w" \
    -trimpath .


FROM scratch

ARG BIN=endpoint
ARG PORT=8080

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /src/${BIN}/build/${BIN} /${BIN}

EXPOSE ${PORT}

USER ${BIN}:${BIN}

ENTRYPOINT [ "/endpoint" ]
