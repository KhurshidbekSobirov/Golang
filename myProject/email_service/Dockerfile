FROM golang:1.17-alpine
RUN mkdir email_service
COPY . /email_service
WORKDIR /email_service
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9002