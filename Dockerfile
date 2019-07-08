FROM golang:latest

EXPOSE 65533

WORKDIR /usr/BurgerCoin/src

#Copy dir
COPY ./ ./