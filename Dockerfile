FROM docker.io/library/golang:1.17-alpine as build

WORKDIR /app
ADD . .

#RUN go mod download
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go build -o /podsdog

#https://hub.docker.com/_/scratch
FROM docker.io/library/scratch
#FROM docker.io/library/busybox
WORKDIR /app

COPY --from=build /podsdog ./podsdog

EXPOSE 8080
VOLUME /

CMD [ "/app/podsdog" ]
