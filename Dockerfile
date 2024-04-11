FROM registry.docker.ir/golang:1.22
WORKDIR app/quizzer
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build
CMD ["bin/main"]
EXPOSE 8080
