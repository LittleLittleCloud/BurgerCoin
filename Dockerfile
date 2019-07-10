FROM golang:latest

EXPOSE 65533

WORKDIR /usr/BurgerCoin/src

#Run CMD
#Config Git
git config --global user.email "g2260578356@gmail.com"
git config --global user.name "BigMiao"

#Copy dir
COPY ./ ./