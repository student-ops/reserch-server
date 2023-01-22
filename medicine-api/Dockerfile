FROM node:alpine

WORKDIR /usr/app
RUN mkdir -p ./cmd/api
RUN mkdir ./prisma

COPY ./cmd/api ./cmd/api
COPY *.json ./
COPY ./prisma /usr/app/prisma
RUN npm install
## warning
RUN npx prisma generate

CMD ["npm", "run","start:migrate:prod"]