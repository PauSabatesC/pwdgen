FROM golang:1.18.1-buster AS build

ENV GO111MODULE=auto

ARG version

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go mod tidy
RUN go build -ldflags "-X github.com/PauSabatesC/pwdgen/cmd.versionTag=$version" -o /pwdgen

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /pwdgen /pwdgen

USER nonroot:nonroot

ENTRYPOINT ["/pwdgen"]
