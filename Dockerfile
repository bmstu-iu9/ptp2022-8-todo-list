# syntax=docker/dockerfile:1

FROM node:12.22.12 AS base
ENV CHOKIDAR_USEPOLLING=1

WORKDIR /code
ADD package.json /code
ADD yarn.lock /code

RUN yarn install

COPY frontend/ /code/frontend/

FROM base AS lint
ADD .prettierrc /code
CMD ["npm", "run", "lint"]

FROM base AS default
CMD ["npm", "run", "dev"]
EXPOSE 8000

FROM base AS deploy
CMD [ "npm", "run", "deploy" ]
EXPOSE 3000