# A test docker app called cors.
FROM ray-xyz.com:9090/ubuntu1604

# Author
MAINTAINER Ray Wang => rwthecoder@gmail.com

# Say something
RUN echo "This is a app from docker, it will run under kubernetes."

RUN mkdir /soft

ADD cors /soft

EXPOSE 8090

CMD ["--port 8090"]

ENTRYPOINT /soft/cors

