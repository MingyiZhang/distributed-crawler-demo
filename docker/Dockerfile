FROM golang:alpine AS builder

# build app in path
WORKDIR /app
COPY . .
RUN go get ./...
ARG path
RUN go build -o app $path

# get wait-for-it
WORKDIR /
RUN apk add --no-cache git && \
    git clone https://github.com/vishnubob/wait-for-it.git

#
FROM golang:alpine
WORKDIR /
RUN apk add --no-cache bash
COPY --from=builder /app/app /wait-for-it/wait-for-it.sh /bin/
CMD ["app"]