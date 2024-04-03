FROM node:21.6.2-alpine
WORKDIR /app
COPY package.json .
RUN npm i
COPY . .
EXPOSE 5000
CMD ["npm", "run", "dev"]