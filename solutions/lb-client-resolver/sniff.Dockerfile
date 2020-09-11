# release
FROM alpine as release

WORKDIR /app

COPY  ./sniff /app/sniff
