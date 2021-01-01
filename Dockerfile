FROM golang

RUN apt-get -y update
RUN apt-get install -y tree wget curl

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

USER root

WORKDIR $GOPATH/src/test_avito
ADD ./ $GOPATH/src/test_avito

RUN chmod +x ./build.sh
RUN ./build.sh

CMD ["./server.app"]