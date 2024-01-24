FROM almalinux:9

RUN dnf clean all \
 && dnf -y install pciutils kmod \
 && dnf clean all 

COPY entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]
ENTRYPOINT ["/entrypoint.sh"]
