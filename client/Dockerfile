FROM node:18.17.0-alpine as base

FROM base as deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY package.json /app
COPY yarn.lock /app
RUN --mount=type=secret,id=npmrc,target=/root/.npmrc yarn install

FROM base as builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN yarn build

FROM base AS runner
WORKDIR /app

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public

RUN mkdir .next
RUN chown nextjs:nodejs .next

COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static


USER nextjs

EXPOSE 3000
ENV HOSTNAME "0.0.0.0"

CMD ["node", "server.js"]