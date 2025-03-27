FROM ghcr.io/luanti-org/luanti:5.11.0
USER root
RUN apk update && apk add git && \
    mkdir -p /root/.minetest/games && \
    git clone https://gitlab.com/sztest/nodecore /root/.minetest/games/nodecore