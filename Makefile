.PHONY: list_topics drivers connectors consumer connect_plugins data

list_topics:
	docker-compose run --rm kafka kafka-topics --describe --zookeeper zookeeper:2181

drivers:
	curl -k -SL "https://dev.mysql.com/get/Downloads/Connector-J/mysql-connector-java-5.1.39.tar.gz" | tar -xzf - -C connect/jars --strip-components=1 mysql-connector-java-5.1.39/mysql-connector-java-5.1.39-bin.jar

connectors:
	./bin/postConnectors.sh

# Avro console consumer to log out events in variants cdc topic
consumer:
	docker-compose exec schema-registry kafka-avro-console-consumer --from-beginning --bootstrap-server kafka:29092 --topic variants 

connect_plugins:
	tar -xvf ./plugins/confluentinc-kafka-connect-elasticsearch-10.0.2.zip -C ./tmp/jars 
	tar -xvf ./plugins/jcustenborder-kafka-connect-spooldir-2.0.46.zip -C ./tmp/jars 

data:
	unzip ./data/variants-fixed.tsv.zip -d ./variants_svc/data
