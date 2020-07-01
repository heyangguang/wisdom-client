# VERSION 1.0
# Author: heyang

FROM 172.16.140.21/heyang/busybox:1.28.4-glibc

MAINTAINER heyang <13833232533@163.com>

COPY wisdomc-ctl /bin/wisdomc-ctl

RUN chmod +x /bin/wisdomc-ctl

ARG CONFIG

ENV config_env_var=$CONFIG

RUN mkdir -p /opt/wisdom

WORKDIR /opt/wisdom

RUN mkdir logs && mkdir conf

COPY conf/tj-config.yaml /opt/wisdom/conf/tj-config.yaml
COPY conf/nj-config.yaml /opt/wisdom/conf/nj-config.yaml
COPY conf/hz-config.yaml /opt/wisdom/conf/hz-config.yaml

CMD /bin/wisdoms-ctl run --config "$config_env_var"