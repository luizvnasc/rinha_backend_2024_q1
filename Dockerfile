FROM golang:latest as build

ADD . /src

#disable crosscompiling 
ENV cgo_enabled=0

#compile linux only
ENV goos=linux

#build the binary with debug information removed
RUN go build -C /src/cmd -ldflags '-w -s' -a -installsuffix cgo -o /bin/rinha  


FROM alpine:latest

ENV PORT=3000
ENV POSTGRES_DNS=""

RUN apk add libc6-compat

COPY --from=build /bin/rinha /bin/rinha

RUN ls /bin

EXPOSE ${PORT}

CMD [ "/bin/rinha" ] 