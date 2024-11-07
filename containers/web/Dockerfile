FROM node:22-alpine

WORKDIR /app

COPY ./web .

RUN cp .env .env.local
RUN npm install && npm run build

EXPOSE 3000

RUN chmod +x entrypoint.sh
CMD ["sh", "entrypoint.sh"]
