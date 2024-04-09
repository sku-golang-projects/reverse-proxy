FROM node:slim

WORKDIR /home/app

COPY package.json package-lock.json index.js /home/app/

RUN npm install

CMD ["npm", "run", "start"]