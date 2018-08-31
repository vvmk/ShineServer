#BUILD-GO
FROM golang:alpine AS build-go
MAINTAINER Vincent Masiello <vincentmasiello@gmail.com>
RUN apk --no-cache add git bzr mercurial
ENV D=/go/src/github.com/vvmk/shineserver
RUN go get -u github.com/golang/dep/...
ADD ./Gopkg.* $D/
RUN cd $D && dep ensure -v --vendor-only
ADD . $D
RUN cd $D && go build -o shineserver && cp shineserver /tmp/

#FINAL
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-go /tmp/shineserver /app/
EXPOSE 8001
ENTRYPOINT ["./shineserver"]
