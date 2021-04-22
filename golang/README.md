# Docker / Go言語

### サーバー用のコード

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
}
```

### Dockerfile

```docker
# ベースとなるDockerイメージ指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /go/src/app

# コンテナ起動時にアプリケーションを起動
CMD ["go", "run", "main.go"]
```

### docker-compose.yml

```yaml
version: '3'
services:
  app:
    build: .
    tty: true # コンテナの起動永続化
    ports:
      - 8080:8080
    volumes:
      - ../:/go/src/app # マウントディレクトリ指定
```

### 起動

```bash
docker-compose build
docker-compose up -d
```