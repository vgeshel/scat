FROM scratch
MAINTAINER vgeshel@gmail.com
ADD scat /bin/scat
CMD ["/bin/scat"]