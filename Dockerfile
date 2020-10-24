FROM golang:1.15-alpine
WORKDIR /src
COPY . /src
RUN apk add git
RUN ./build /bin/go-starter
ENTRYPOINT ["/bin/go-starter"]
