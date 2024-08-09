FROM node:lts-alpine
ENV NODE_ENV=production

WORKDIR /usr/src/app
COPY ["package.json", "yarn.lock", "./"]
RUN yarn install --frozen-lockfile
COPY . .
RUN chown -R node /usr/src/app

EXPOSE 3000

USER node
CMD ["yarn", "start"]
