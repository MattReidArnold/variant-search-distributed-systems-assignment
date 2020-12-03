#!/bin/bash

CONNECT_HOST='localhost'
CONNECT_PORT='8083'
CONNECTORS_URL="http://$CONNECT_HOST:$CONNECT_PORT/connectors"

curl -X DELETE "$CONNECTORS_URL/avro-file-sink"
curl -X DELETE "$CONNECTORS_URL/postgresql-sink"
curl -X DELETE "$CONNECTORS_URL/elasticsearch-sink"

sleep 3

curl -X POST $CONNECTORS_URL -H "Content-Type: application/json" --data '{
  "name": "avro-file-sink", 
  "config": {
    "connector.class":"org.apache.kafka.connect.file.FileStreamSinkConnector", 
    "tasks.max":"1", 
    "topics":"variants", 
    "file": "/tmp/filesink/avro-file-sink-variants.txt"
  }
}'

curl -X POST $CONNECTORS_URL -H "Content-Type: application/json" -d '{
 "name": "postgresql-sink",
 "config": {
   "connector.class": "io.confluent.connect.jdbc.JdbcSinkConnector",
   "connection.url": "jdbc:postgresql://postgres:5432/variant-search",
   "connection.user": "admin",
   "connection.password": "password",
   "tasks.max": "1",
   "topics": "variants",
   "auto.create": true
 }
}'

curl -X POST $CONNECTORS_URL -H "Content-Type: application/json" -d '{
 "name": "elasticsearch-sink",
 "config": {
   "connector.class": "io.confluent.connect.elasticsearch.ElasticsearchSinkConnector",
   "connection.url": "http://elasticsearch:9200",
   "tasks.max": "1",
   "topics": "variants",
   "type.name": "_doc"
 }
}'
