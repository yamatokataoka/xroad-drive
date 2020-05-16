FROM node:12.16.3-alpine AS develop
ENV YARN_CACHE_FOLDER /xroad-drive-ui/.yarn
COPY package.json yarn.lock /xroad-drive-ui/
WORKDIR /xroad-drive-ui
RUN yarn install
COPY . /xroad-drive-ui

FROM develop AS build
RUN yarn run build

FROM nginx:stable-alpine AS production
RUN rm /etc/nginx/conf.d/*
RUN rm /usr/share/nginx/html/*
COPY --from=build /xroad-drive-ui/dist /opt/xroad-drive/xroad-drive-ui
COPY docker/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]