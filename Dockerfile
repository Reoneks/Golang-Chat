FROM golang:latest
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN go get \
    && go build -o app . \
    && chmod +x app \
    && mv ./app /usr/local/bin/app

ENTRYPOINT ["app"]