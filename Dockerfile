FROM golang:1.8
COPY PackageIndexer /PackageIndexer
WORKDIR /PackageIndexer
EXPOSE 8080
RUN go build
CMD ["PackageIndexer"]