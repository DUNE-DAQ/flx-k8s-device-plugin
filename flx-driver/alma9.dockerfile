FROM almalinux:9

ARG HTTP_PROXY=local
ARG HTTPS_PROXY=local
ENV HTTP_PROXY=${HTTP_PROXY}
ENV HTTPS_PROXY=${HTTPS_PROXY}

RUN yum clean all \
 && yum -y install pciutils \
 && yum clean all 

COPY entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]
ENTRYPOINT ["/entrypoint.sh"]
