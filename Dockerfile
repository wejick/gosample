FROM scratch
LABEL Name=gosample Version=0.0.1
#SET ENV Variable here
#ENV TKPENV production
ADD files/ /
ADD gosample /
EXPOSE 9000
CMD [ "/gosample" ]