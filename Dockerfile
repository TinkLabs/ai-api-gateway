FROM bananabb/ubuntu-go:go.1.3.1

# Copy project to container
ADD . /var/src

# Setup project
RUN cd /var/src/gateway \
 dep ensure 

# Init CAT env
RUN mkdir -p /data/appdatas/cat
RUN mkdir -p /data/applogs/cat
RUN echo '<?xml version="1.0" encoding="utf-8"?>' >> /data/appdatas/cat/client.xml 
RUN echo '<config xmlns:xsi="http://www.w3.org/2001/XMLSchema" xsi:noNamespaceSchemaLocation="config.xsd">' >> /data/appdatas/cat/client.xml 
RUN echo '    <servers>' >> /data/appdatas/cat/client.xml 
RUN echo '        <server ip="10.1.0.172" port="2280" http-port="8080" />' >> /data/appdatas/cat/client.xml 
RUN echo '    </servers>' >> /data/appdatas/cat/client.xml 
RUN echo '</config>' >> /data/appdatas/cat/client.xml

# Basic setup
WORKDIR /var/src/gateway
EXPOSE 80 443

CMD ["bash"]