FROM alpine:3.17

ARG RELEASE=latest

WORKDIR /app

RUN apk add curl tar gunzip && \
  curl https://github.com/GIT_OWNER/GIT_PROJECT/releases/download/${RELEASE}/GIT_PROJECT_${RELEASE}_linux_amd64.tar.gz | tar -xzvf -

USER nonroot:nonroot

ENTRYPOINT ["/app/GIT_PROJECT", "sample-command"]
