FROM alpine
WORKDIR /logtest
COPY app app
ENTRYPOINT ["/logtest/app"]

