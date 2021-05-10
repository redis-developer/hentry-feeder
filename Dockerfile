FROM frolvlad/alpine-gxx

# labels
LABEL maintainer="Yash Kumar Verma yk.verma2000@gmail.com"

# set environment varialbes
ENV PORT=80
ENV REDIS_URL=localhost:6379

# copy project to working directory
WORKDIR /

# take built packge into container
COPY ./build/hentry-feeder ./hentry-feeder

# grant permissions to execute
RUN chmod +x hentry-feeder

# launch server
CMD ["./hentry-feeder"]

# exporse listening port
EXPOSE 80
