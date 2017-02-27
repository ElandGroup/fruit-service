FROM jaehue/golang-onbuild
MAINTAINER jang.jaehue@eland.co.kr

# install go packages

# add application
WORKDIR /go/src/fruit-service
ADD . /go/src/fruit-service
RUN go install

EXPOSE 5000

CMD ["/go/bin/fruit"]