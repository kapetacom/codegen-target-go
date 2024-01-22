FROM node:20 as build

ADD package.json /application/package.json
ADD package-lock.json /application/package-lock.json

WORKDIR /application

RUN npm install

ADD db /application/db
ADD tsconfig.json /application/tsconfig.json
ADD kapeta.yml /application/kapeta.yml
ADD index.ts /application/index.ts
ADD src /application/src

RUN npm run build

FROM node:20-alpine3.18
COPY --from=build /application /application
WORKDIR /application
RUN npm run migrate:generate --if-present #Needed for Prisma to use the right binaries
HEALTHCHECK --interval=5s --timeout=15s --start-period=5s --retries=10 CMD wget --no-verbose --tries=1 --spider "http://localhost:80/.kapeta/health" || exit 1

ENTRYPOINT [ "npm", "start" ]
