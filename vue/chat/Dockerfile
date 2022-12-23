FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm i --legacy-peer-deps
COPY . .
# RUN npm run build