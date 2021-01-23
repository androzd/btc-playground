FROM ubuntu:20.04

RUN apt-get update

RUN apt-get -y install curl

RUN curl -OL https://bitcoin.org/bin/bitcoin-core-0.20.0/bitcoin-0.20.0-x86_64-linux-gnu.tar.gz

RUN tar zxvf bitcoin-0.20.0-x86_64-linux-gnu.tar.gz

RUN ln -s /bitcoin-0.20.0/bin/bitcoin-cli /bitcoin-cli

# rpc
EXPOSE 18444/tcp
# p2p
EXPOSE 18443/tcp

ENTRYPOINT ["/bitcoin-0.20.0/bin/bitcoind"]
