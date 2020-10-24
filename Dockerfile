FROM golang:1.15-alpine
WORKDIR /src
COPY . /src
RUN apk add git
RUN ./build /bin/snarecc-pages
ENTRYPOINT ["/bin/snarecc-pages"]
