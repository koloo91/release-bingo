FROM node:12-alpine as nodeBuilder
WORKDIR /builder
COPY frontend/ .
RUN npm config set unsafe-perm true
RUN npm install -g @angular/cli
RUN npm install && npm run build

FROM golang:1.14-alpine AS goBuilder
WORKDIR /builder
ADD backend/ .
RUN go build -o release-bingo

FROM alpine
WORKDIR /app
COPY --from=goBuilder /builder/release-bingo .
COPY --from=nodeBuilder /builder/dist/release-bingo/ public/
COPY backend/migrations/ migrations/
ENTRYPOINT ["./release-bingo"]
