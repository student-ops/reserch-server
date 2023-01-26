FROM node:alpine

WORKDIR /usr/app
RUN mkdir ./prisma

COPY ./index.ts ./index.ts
COPY *.json ./
COPY ./prisma ./prisma
RUN npm install
## warning
RUN npx prisma generate

# CMD ["npm", "run","start"]