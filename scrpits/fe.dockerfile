# pull official base image
FROM node:16.15.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

#RUN npm install -g create-react-app@18.1.0
# install app dependencies
#COPY /fe/fe-table-task/package.json ./
#COPY /fe/fe-table-task/package-lock.json ./

# add app
COPY /fe/fe-table-task /app

RUN npm install

# start app
CMD ["npm", "start"]