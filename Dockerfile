FROM golang:1.17-buster
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
COPY . /asset
RUN pwd
RUN go mod download
COPY . .
RUN go mod tidy
RUN go install
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["running-backend"]
EXPOSE 8000