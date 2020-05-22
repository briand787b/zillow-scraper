FROM node:14.3
WORKDIR /app
COPY package*.json ./
RUN npm install
EXPOSE 8000
COPY . ./
CMD [ "npm", "run", "dev" ]
