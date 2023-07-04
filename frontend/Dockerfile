# ビルド環境
FROM node:18.16.0 as build
WORKDIR /app
COPY package.json ./
RUN yarn install
COPY . ./
RUN yarn build

# 本番環境
FROM nginx:1.19.0-alpine as production
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
