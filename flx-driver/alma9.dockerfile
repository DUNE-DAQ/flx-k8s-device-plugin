FROM almalinux:9

RUN yum clean all \
 && yum -y install pciutils \
 && yum clean all 

COPY entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]
ENTRYPOINT ["/entrypoint.sh"]
