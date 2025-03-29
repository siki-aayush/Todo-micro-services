FROM oven/bun:latest

RUN mkdir /app

WORKDIR /app

COPY . .

RUN bun install

CMD ["bun", "dev", "--host", "0.0.0.0", "--port", "5173"]
