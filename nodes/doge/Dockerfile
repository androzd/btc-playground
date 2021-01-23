FROM ubuntu:20.04

RUN apt-get update

RUN apt-get -y install curl

RUN curl -OL https://github.com/dogecoin/dogecoin/releases/download/v1.14.2/dogecoin-1.14.2-x86_64-linux-gnu.tar.gz

RUN tar zxvf dogecoin-1.14.2-x86_64-linux-gnu.tar.gz

RUN ln -s /dogecoin-1.14.2/bin/dogecoin-cli /dogecoin-cli

# rpc
EXPOSE 18444/tcp
# p2p
EXPOSE 18443/tcp

ENTRYPOINT ["/dogecoin-1.14.2/bin/dogecoind"]
