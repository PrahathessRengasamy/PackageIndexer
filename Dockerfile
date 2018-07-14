FROM golang:1.8
COPY ./Solution /Solution
WORKDIR /Solution
RUN go build
CMD ["./Solution"]
EXPOSE 8080