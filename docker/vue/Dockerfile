FROM nginx AS nginx
WORKDIR /chatait
RUN mkdir /chatait/backendVue
RUN mkdir /chatait/backendVue/dist
RUN mkdir /chatait/frontendVue
RUN mkdir /chatait/frontendVue/dist
COPY ./backendVue/dist /chatait/backendVue/dist/
COPY ./frontendVue/dist /chatait/frontendVue/dist/
COPY ./nginx.conf /etc/nginx/conf.d/chatait.conf







 

