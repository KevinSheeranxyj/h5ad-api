## 构建
FROM easyunit/go:1.15 as builder

ARG Dir=/go/src/robot

WORKDIR ${Dir}

COPY . .

RUN go build -o robot

## 运行环境
FROM easyunit/alpine:3-tzdata

ARG Dir=/go/src/robot

WORKDIR ${Dir}

COPY ./logs ./logs

COPY --from=builder  ${Dir}/robot ./
COPY --from=builder  ${Dir}/.env ./

EXPOSE 8000
CMD ["./robot"]