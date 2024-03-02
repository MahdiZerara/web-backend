FROM golang:1.21.3-bookworm@sha256:3ce8313c3513515040870c55e0c041a2b94f3576a58cfd3948633604214aa811
LABEL maintainer="Mahdi Zerara - Web Backend"

ENV REST_SERVER_PORT="3000"

RUN apt-get update && apt-get upgrade -y

RUN mkdir /app
COPY /src /app
WORKDIR /app

RUN go mod tidy
RUN go build -buildvcs=false -o main ./cmd 

EXPOSE $REST_SERVER_PORT

CMD ["./main"]
