FROM scratch
MAINTAINER vgeshel@gmail.com
ADD scat /bin/scat
RUN chmod a+x /bin/scat
CMD ["/bin/scat"]