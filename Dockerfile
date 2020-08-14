FROM debian
COPY bin/go-starter /bin/go-starter
ENTRYPOINT ["/bin/go-starter"]
