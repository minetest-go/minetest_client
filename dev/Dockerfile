FROM ghcr.io/minetest/minetest:5.9.0
USER root
RUN apk update && apk add git && \
    mkdir -p /root/.minetest/games && \
    git clone https://gitlab.com/sztest/nodecore /root/.minetest/games/nodecore