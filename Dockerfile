FROM golang:1.8
COPY ./Solution /Solution
WORKDIR /Solution
EXPOSE 8080
RUN go build
CMD ["Solution"]