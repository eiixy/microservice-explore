.PHONY: elk.%
ELK_VERSION=7.12.0

## elk.install: 安装 ELK 环境
elk.install: elasticsearch.install logstash.install kibana.install

elk.run: elasticsearch.run logstash.run kibana.run

elk.version:
	@echo "elk version: $(ELK_VERSION)"

elasticsearch.install:
	wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$(ELK_VERSION)-linux-x86_64.tar.gz
	tar -xvf elasticsearch-$(ELK_VERSION)-linux-x86_64.tar.gz -C /usr/local/
	mv /usr/local/elasticsearch-$(ELK_VERSION) /usr/local/elasticsearch

logstash.install:
	wget https://artifacts.elastic.co/downloads/logstash/logstash-$(ELK_VERSION)-linux-x86_64.tar.gz
	tar -xvf logstash-$(ELK_VERSION)-linux-x86_64.tar.gz -C /usr/local/
	mv /usr/local/logstash-$(ELK_VERSION) /usr/local/logstash

kibana.install:
	wget https://artifacts.elastic.co/downloads/kibana/kibana-$(ELK_VERSION)-linux-x86_64.tar.gz
	tar -xvf kibana-$(ELK_VERSION)-linux-x86_64.tar.gz -C  /usr/local/
	mv /usr/local/kibana-$(ELK_VERSION)-linux-x86_64 /usr/local/kibana

elasticsearch.run:
	@echo "run elasticsearch"
	/usr/local/elasticsearch/bin/elasticsearch -d

logstash.run:
	@echo "run logstash"
	/usr/local/logstash/bin/logstash -f /usr/local/logstash/config/pv_track.conf

kibana.run:
	@echo "run kibana"
	nohup /usr/local/kibana/bin/kibana &

