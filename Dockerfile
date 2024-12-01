FROM node:22-alpine3.20

WORKDIR /app

COPY server/package.json server/package-lock.json ./

RUN npm install

COPY server/src ./

CMD ["node","index.js"]