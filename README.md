About

This is a pet project to understand kafka better.
It contains a go-client (2 consumers and producer) and a kafka cluster.
2 instances of kafka will be started.
The producer will keep producing events in an infinite loop and there are 2 consumers which are reading from the producer at different rates.

Test cases done via commands and understandings

1. If there are 3 brokers , and we create a topic with replication factor as 3 and 3 partitions. The topic will be distributed across all 3 brokers , i.e each partition of the topic will be copied to other 2 brokers and each partition will have a leader.If a broker was to go down , other brokers will take over and consumers wont be impacted.
2. If we have just 1 partition , but replication factor of 3 , all consumers will read from the same partition at the same time. If the broker goes down , consumers/producers wont be impacted because we have other brokers who can take over and they have the latest data.
3. Consider 3 brokers , and we create a topic with a replication factor of 1 and partition 1 . In this case , a leader will be selected and all events will be stored in the same broker. If this broker was to go down , the consumers will lose messages.

TODO and WIP
Automate creation of topics.
Automate starting of producers and both consumers.

How to use the repo

1. Switch to kafka directory.
2. docker-compose up -d
3. Create a topic by below command, to do this , exec into anyone of the kafka brokers 
   docker exec -it <container id> sh 
   kafka-topics.sh --bootstrap-server broker-1:9092 --create --topic my-topic --partitions 3 --replication-factor 3
4. go run main.go in the producer directory.
5. go run main.go in consumer directory.
6. go run main.go in consumer-2 directory.

Yes , naming conventions are bad , but this is just to help me understand kafka better.


Kafka commands

Create topic
kafka-topics.sh --bootstrap-server broker-1:9092 --create --topic my-topic --partitions 3 --replication-factor 3

Describe topic
kafka-topics.sh --bootstrap-server broker-1:9092 --describe --topic my-topic


Output
Topic: my-topic	TopicId: 14qk5dr4S_SQT8bZRFVS4w	PartitionCount: 3	ReplicationFactor: 3	Configs: segment.bytes=1073741824
	Topic: my-topic	Partition: 0	Leader: 1	Replicas: 1,2,3	Isr: 1,2,3
	Topic: my-topic	Partition: 1	Leader: 2	Replicas: 2,3,1	Isr: 2,3,1
	Topic: my-topic	Partition: 2	Leader: 3	Replicas: 3,1,2	Isr: 3,1,2


Kafka concepts
Kafka is a distributed event store and stream platform.
Kafka is highly reliable , scalable and low-latency platform due to some concepts we will see below.

Topic
A kafka topic is used to organize messages/events.
Example , a topic to store scores of a live updates of football game can be named “germany-vs-portugal” .
We can now push all live updates into this topic relating to the mentioned game.
Similarly a cluster will have multiple topics with producers constantly pushing data/events into these topics and consumers reading from relevant topics.

The command to create topic is 
kafka-topics.sh --bootstrap-server broker-1:9092 --create --topic my-topic --partitions 3 --replication-factor 3

If we try to understand some arguments passed here , we can understand what makes kafka highly reliable and scalable for producers/consumers.

--partitions 3
The above command means create 3 partitions of the given topic.
Partition 0 , 1 and 2.
Now when a producer pushes a message into the topic named “my-topic” , the produced message will land into one of the partitions.
It depends whether a key is given to the message.
If a key is given in the message , the message will be assigned to a partition based on calculations done on the key.
If key is not provided , the messages will be assigned to partitions in a round robin manner.
And it is this property of kafka that can help process messages faster at the consumer side.

If the number of consumers (of the same service) is same as number of partitions , every instance of that microservice will read from a partition.
Eg - If a micro service was deployed in kubernetes with 3 replicas , each replica would consume from one partition , and this would make processing fast.


--replication-factor 3
This makes kafka highly reliable.
This means the topic will be replicated thrice across 3 broker.
Note - Replication factor can be as big as number of brokers. If we try to create a topic with more replication factor than number of brokers , we get an error.
This means , the above command will create 3 copies of “”my-topic” topic and will stores across 3 brokers.
This also means every partition will be replicated.
So 3 partitions X 3 times divided across 3 brokers.

Every partition has a leader assigned. This means reading/writing from a particular partition will only happen through the leader.
The message which is stored in the given partition will then be replicated across other brokers.
If no replication factor is given , if the broker goes down , messages will be lost.


With replication factor , even if a broker goes down , as messages are replicated across brokers , a new leader for a partition will be selected and there wont be loss of messages.


Key observations-
Even if consumer loses connection to the broker , the consumer will read from the last read message. (feature comes out of the library used , it commits offset of the last read message after read is complete.)
For an existing consumer , offset as newest or oldest does not make a difference , in both cases , the consumer will read from the last read message.
But this offset (oldest or newest ) is important for adding a new consumer to an existing topic.
If newest , it wont read old messages.(the consumer will start reading from the latest message)
If oldest , it will read all old messages starting from offset 0 .






























































































