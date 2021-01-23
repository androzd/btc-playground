FROM ubuntu:20.04

RUN apt-get update

RUN apt-get -y install curl

RUN curl -OL https://download.litecoin.org/litecoin-0.18.1/linux/litecoin-0.18.1-x86_64-linux-gnu.tar.gz

RUN tar zxvf litecoin-0.18.1-x86_64-linux-gnu.tar.gz

RUN ln -s /litecoin-0.18.1/bin/litecoin-cli /litecoin-cli

# rpc
EXPOSE 18444/tcp
# p2p
EXPOSE 18443/tcp

ENTRYPOINT ["/litecoin-0.18.1/bin/litecoind"]
