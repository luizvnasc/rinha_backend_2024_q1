FROM golang:latest as build

ADD . /src

#disable crosscompiling 
ENV cgo_enabled=0

#compile linux only
ENV goos=linux

#build the binary with debug information removed
RUN go build -C /src/cmd -ldflags '-w -s' -a -installsuffix cgo -o /bin/rinha  

RUN ls /bin

FROM scratch

ENV PORT=3000
ENV POSTGRES_DNS=

COPY --from=build /bin/rinha /

EXPOSE ${PORT}

ENTRYPOINT [ "rinha" ] 