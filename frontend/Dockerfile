# syntax=docker/dockerfile:1

FROM node:12.18.1

WORKDIR /app

COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm install --unsafe-perms -g

COPY . /app

CMD npm run dev
EXPOSE 80
