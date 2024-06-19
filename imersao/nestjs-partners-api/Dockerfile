FROM node:21-slim

RUN apt update && apt install -y openssl procps

RUN npm install -g @nestjs/cli@10.3.2

RUN npm i -g npm@10.8.1

WORKDIR /home/node/app

USER node

CMD tail -f /dev/null