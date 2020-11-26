# Variant Search Distributed Systems Assignment

## Assignment

Treating `data/variants.tsv.zip` as source data:
1. Write a `Kafka` producer which publishes records onto a `Kafka` topic
2. Come with a strategy to populate published records into an `Elasticsearch` index
3. Come with a strategy to populate published records into a `postgres` db table [Optional]

We're trying to mimic a 1 source -> 2 sink scenario, but you can choose to treat #3 as optional. 
Please craft solutions prioritizing maintainability and best practices.           

## Bootstrap

- A `docker-compose` setup is provided a starting point, which sets up `Kafka`, `Zookeeper`, `Elasticsearch`, `Kibana` & `Postgres`
- Feel free to use a more compact setup which reduces # of individual pieces (eg: `lenses.io`)  

## Implementation

`Python` or a `JVM` language (`Kotlin`, `Scala`, `Java`) is preferred, but feel free to employ any programming language  
