# syntax=docker/dockerfile:1

FROM node:12.22.12 AS base
ENV CHOKIDAR_USEPOLLING=1

RUN npm install -g parcel

ADD . /code
WORKDIR /code

FROM base AS lint
CMD npm run lint

FROM base
CMD npm run dev
EXPOSE 80
