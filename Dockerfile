FROM ubuntu

WORKDIR /usr/home

ADD main mian

RUN ./main