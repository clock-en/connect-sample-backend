FROM postgres:17.3

# pgvector拡張機能をインストール
RUN apt-get update && \
    apt-get install -y git make gcc postgresql-server-dev-17 && \
    cd /tmp && \
    git clone --branch v0.8.0 https://github.com/pgvector/pgvector.git && \
    cd pgvector && \
    make && \
    make install && \
    cd ../ && rm -rf pgvector

COPY .environments/local/docker/db/initdb.d /docker-entrypoint-initdb.d
