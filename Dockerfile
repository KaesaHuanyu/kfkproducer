FROM daocloud.io/golang
maintainer kaesa.li@daocloud.io
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go install kafkaproducer
ENV TZ Asia/Shanghai
entrypoint ["/gopath/app/bin/kafkaproducer"]