# syntax=docker/dockerfile:1

FROM node:12.22.12
ENV CHOKIDAR_USEPOLLING=1

RUN npm install -g parcel

ADD . /code
WORKDIR /code

CMD npm run dev
EXPOSE 80
