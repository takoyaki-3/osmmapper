FROM centos:centos8.1.1911
MAINTAINER Admin <admin@admin.com>

# "RUN" do at docker build
RUN dnf install -y wget
ADD ./start.sh /start.sh

CMD ./start.sh