FROM node:22-alpine

WORKDIR /app

COPY ./web .

ARG EXTERNAL_API_URL

RUN cp .env .env.local
RUN sed -i "s|^NEXT_PUBLIC_API_URL=.*|NEXT_PUBLIC_API_URL=$EXTERNAL_API_URL|" .env.local

RUN npm install && npm run build

EXPOSE 3000

CMD ["npm", "start"]
