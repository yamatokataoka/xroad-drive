FROM node:lts-alpine as builder
ENV YARN_CACHE_FOLDER /fileservice-ui/.yarn
COPY package.json yarn.lock /fileservice-ui/
WORKDIR /fileservice-ui
RUN yarn install
COPY . /fileservice-ui
RUN yarn run build

FROM nginx:stable-alpine AS production
RUN rm /etc/nginx/conf.d/*
RUN rm /usr/share/nginx/html/*
COPY --from=builder /fileservice-ui/dist /opt/fileservice/fileservice-ui
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]