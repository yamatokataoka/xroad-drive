FROM node:12.16.3-alpine AS develop
ENV YARN_CACHE_FOLDER /fileservice-ui/.yarn
COPY package.json yarn.lock /fileservice-ui/
WORKDIR /fileservice-ui
RUN yarn install
COPY . /fileservice-ui

FROM develop AS build
RUN yarn run build

FROM nginx:stable-alpine AS production
RUN rm /etc/nginx/conf.d/*
RUN rm /usr/share/nginx/html/*
COPY --from=build /fileservice-ui/dist /opt/fileservice/fileservice-ui
COPY docker/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]