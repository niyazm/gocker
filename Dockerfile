FROM alpine:3.13.5

LABEL maintainer "Niyaz Mahmud <niyaztron@gmail.com>"

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /go/src/github.com/niyazm

COPY . .

RUN go install

EXPOSE 8080

#
#

ENTRYPOINT ["dockerhw"]


