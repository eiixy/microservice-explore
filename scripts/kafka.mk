.PHONY: kafka.install
KAFKA_VERSION=2.12-2.8.0

kafka.install:
	wget https://mirrors.bfsu.edu.cn/apache/kafka/2.8.0/kafka_${KAFKA_VERSION}.tgz
	tar -xvf kafka_${KAFKA_VERSION}.tgz
	sudo mv kafka_${KAFKA_VERSION} /usr/local/kafka

kafka.run:
	@echo "remove logs dir"
	rm -rf /tmp/kafka-logs/* /tmp/zookeeper/*
	@echo "run zookeeper"
	/usr/local/kafka/bin/zookeeper-server-start.sh -daemon /usr/local/kafka/config/zookeeper.properties &
	@echo "run kafka"
	/usr/local/kafka/bin/kafka-server-start.sh -daemon /usr/local/kafka/config/server.properties &
