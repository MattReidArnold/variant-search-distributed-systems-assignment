#!/bin/bash

CONNECT_HOST='localhost'
CONNECT_PORT='8083'

# curl -X DELETE "http://$CONNECT_HOST:$CONNECT_PORT/connectors/postgres-jdbc-source"
# curl -X DELETE "http://$CONNECT_HOST:$CONNECT_PORT/connectors/tsv-spooldir-source"
curl -X DELETE "http://$CONNECT_HOST:$CONNECT_PORT/connectors/avro-file-sink"

sleep 3

# curl -X POST \
#   -H "Content-Type: application/json" \
#   --data '{ "name": "postgres-jdbc-source", "config": { "connector.class": "io.confluent.connect.jdbc.JdbcSourceConnector", "tasks.max": 1, "connection.url": "jdbc:postgresql://postgres:5432/variant-search", "connection.user": "admin","connection.password": "password", "mode": "incrementing", "incrementing.column.name": "id", "timestamp.column.name": "modified", "topic.prefix": "postgres-jdbc-", "poll.interval.ms": 1000 } }' \
#   "http://$CONNECT_HOST:$CONNECT_PORT/connectors"

curl -X POST -H "Content-Type: application/json" \
  --data '{"name": "avro-file-sink", "config": {"connector.class":"org.apache.kafka.connect.file.FileStreamSinkConnector", "tasks.max":"1", "topics":"variants", "file": "/tmp/filesink/avro-file-sink-variants.txt"}}' \
  "http://$CONNECT_HOST:$CONNECT_PORT/connectors"

# curl -X POST -H "Content-Type: application/json" \
#   --data '{"name": "tsv-spooldir-source", "config": { "tasks.max": 1, "connector.class": "com.github.jcustenborder.kafka.connect.spooldir.SpoolDirCsvSourceConnector", "input.path": "/tmp/input", "input.file.pattern": "variants-new.tsv", "error.path": "/tmp/error", "finished.path": "/tmp/finished", "halt.on.error": false, "topic": "spooldir-tsv-variants", "schema.generation.enabled": true, "csv.first.row.as.header": true, "csv.separator.char": 9 }}' \
#   "http://$CONNECT_HOST:$CONNECT_PORT/connectors"



# curl -X POST -H "Content-Type: application/json" \
#   --data '{"name": "tsv-spooldir-source", "config": { "tasks.max": 1, "connector.class": "com.github.jcustenborder.kafka.connect.spooldir.SpoolDirCsvSourceConnector", "input.path": "/tmp/input", "input.file.pattern": "variants-new.tsv", "error.path": "/tmp/error", "finished.path": "/tmp/finished", "halt.on.error": false, "topic": "spooldir-tsv-variants", "csv.first.row.as.header": true, "csv.separator.char": 9, "key.schema":"{\"type\":\"record\",\"name\":\"VariantKey\",\"namespace\":\"com.github.mattreidarnold.variant\",\"fields\":[],\"connect.name\":\"com.github.mattreidarnold.variant.Key\"}", "value.schema" : "{\"type\":\"record\",\"name\":\"VariantValue\",\"namespace\":\"com.github.mattreidarnold.variant\",\"fields\":[{\"name\":\"Gene\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Nucleotide_Change\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Protein_Change\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Other_Mappings\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Alias\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Transcripts\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Region\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Reported_Classification\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Inferred_Classification\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Source\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Last_Evaluated\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Last_Updated\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"URL\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Submitter_Comment\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Assembly\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Chr\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Genomic_Start\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Genomic_Stop\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Ref\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Alt\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Accession\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Reported_Ref\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Reported_Alt\",\"type\":[\"null\",\"string\"],\"default\":null}],\"connect.name\":\"com.github.mattreidarnold.variant.Value\"}" }}' \
#   "http://$CONNECT_HOST:$CONNECT_PORT/connectors"

# curl -X POST -H "Content-Type: application/json" \
#   --data '{"name": "tsv-spooldir-source", "config": { "tasks.max": 1, "connector.class": "com.github.jcustenborder.kafka.connect.spooldir.SpoolDirCsvSourceConnector", "input.path": "/tmp/input", "input.file.pattern": "variants-new.tsv", "error.path": "/tmp/error", "finished.path": "/tmp/finished", "halt.on.error": false, "topic": "spooldir-tsv-variants", "csv.first.row.as.header": true, "csv.separator.char": 9, "key.schema":"{\"name\":\"com.github.mattreidarnold.variant.VariantKey\",\"type\":\"STRUCT\",\"isOptional\":false,\"fieldSchemas\":{\"Gene\":{\"type\":\"STRING\",\"isOptional\":false}}}", "value.schema" : "{\"name\":\"com.example.users.User\",\"type\":\"STRUCT\",\"isOptional\":false,\"fieldSchemas\":{\"Gene\":{\"type\":[null,\"STRING\"],\"default\":null},\"Nucleotide_Change\":{\"type\":[null,\"STRING\"],\"default\":null}}}" }}' \
#   "http://$CONNECT_HOST:$CONNECT_PORT/connectors"
