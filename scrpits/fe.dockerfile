# pull official base image
FROM node:16.15.0-alpine

# set working directory
WORKDIR /usr/src/fe-task

# add `/app/node_modules/.bin` to $PATH
#ENV PATH /fe-task/node_modules/.bin:$PATH

#RUN npm install -g create-react-app@18.1.0
# install app dependencies
#COPY /fe/fe-table-task/package.json ./
#COPY /fe/fe-table-task/package-lock.json ./

# add app
COPY /fe/fe-table-task  /usr/src/fe-task

RUN npm install

# start app
CMD ["npm", "start"]