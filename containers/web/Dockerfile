FROM node:22-alpine

WORKDIR /app

COPY ./web .
WORKDIR /app/web

RUN npm install && npm run build

EXPOSE 3000

CMD ["npm", "start"]
