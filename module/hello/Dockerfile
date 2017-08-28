# A test docker app called hello.
FROM ubuntu:16.04

# Author
MAINTAINER Ray Wang => rwthecoder@gmail.com

# Say something
RUN echo "This is a app from docker, it will run under kubernetes."

RUN mkdir /soft

ADD hello /soft

EXPOSE 8787

CMD ["--port 8787"]

ENTRYPOINT /soft/hello
