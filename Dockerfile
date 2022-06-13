FROM golang:1.16-alpine

WORKDIR /quiz
ADD . /quiz

RUN cd /quiz && go build

EXPOSE 8080
CMD ./quiz
