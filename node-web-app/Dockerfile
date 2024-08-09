FROM node:20.16.0-alpine

WORKDIR /usr/src/app
COPY ["package.json", "yarn.lock", "./"]
RUN yarn install --frozen-lockfile
COPY . .

RUN chown -R node /usr/src/app
USER node
