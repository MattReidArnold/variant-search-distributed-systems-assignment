.PHONY: list_topics connectors consumer plugins data

# Prepare tsv for consumption
data:
	unzip ./data/variants-fixed.tsv.zip -d ./variants_svc/data

# Install kafka connect plugins
plugins:
	mkdir -p ./tmp/jars
	tar -xvf ./plugins/confluentinc-kafka-connect-elasticsearch-10.0.2.zip -C ./tmp/jars 
	tar -xvf ./plugins/confluentinc-kafka-connect-jdbc-10.0.1.zip -C ./tmp/jars 

# Delete and recreate kafka connect connectors 
connectors:
	./bin/postConnectors.sh

##################################################################################
## Debug
##################################################################################

# List kafka topics
list_topics:
	docker-compose run --rm kafka kafka-topics --describe --zookeeper zookeeper:2181

# Avro console consumer to log out events in variants cdc topic
consumer:
	docker-compose exec schema-registry kafka-avro-console-consumer --from-beginning --bootstrap-server kafka:29092 --topic variants 


