FROM golang:latest


# Configure the repo url so we can configure our work directory:
WORKDIR $GOPATH/src/github.com/vSivarajah/bookstore_users-api

COPY . .

 RUN go get -d -v ./...

 RUN go install -v ./...

# Expose port 8081 to the world:
RUN go build .

EXPOSE 8081



CMD ["./bookstore_users-api"]
