FROM radial/busyboxplus
MAINTAINER vgeshel@gmail.com
ADD scat /bin/scat
RUN /bin/chmod a+x /bin/scat
CMD ["/bin/scat"]