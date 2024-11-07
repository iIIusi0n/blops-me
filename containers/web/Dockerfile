FROM node:22-alpine

WORKDIR /app

COPY ./web .

ARG NEXT_PUBLIC_API_URL=BAKED_NEXT_PUBLIC_API_URL

RUN npm install && npm run build

EXPOSE 3000

CMD ["sh", "replace.sh", "&&", "npm", "start"]
