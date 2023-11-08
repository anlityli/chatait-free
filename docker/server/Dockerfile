FROM golang:1.21

WORKDIR /chatait
RUN mkdir /chatait/config
RUN mkdir /chatait/tmp/
RUN mkdir /chatait/files/
RUN mkdir /chatait/backendServer
RUN mkdir /chatait/backendServer/runtime
RUN mkdir /chatait/frontendServer
RUN mkdir /chatait/frontendServer/runtime
VOLUME /chatait/config
VOLUME /chatait/tmp/
VOLUME /chatait/files/
COPY ./backendServer/chatait-backend-server /chatait/backendServer/
COPY ./frontendServer/chatait-frontend-server /chatait/frontendServer/
COPY ./docker-run.sh /chatait
RUN chmod +x /chatait/backendServer/chatait-backend-server
RUN chmod +x /chatait/frontendServer/chatait-frontend-server

WORKDIR /chatait
CMD /chatait/docker-run.sh