# Установка и запуск

* качаем последную сборку nsqd(https://github.com/nsqio/nsq/releases/tag/v0.3.8) и кидаем бинари в $PATH
* `go get github.com/antonikonovalov/tests-nsqd` - клоникуем репу и ставим бинарь для сравнения логов nsq_to_file
* `go get github.com/ddollar/forego`
* `cd $GOPATH/src/github.com/antonikonovalov/tests-nsqd`
* `forego start -f Procfile.nsqd` - запускаем nsqd_1,nsqd_2,nsqlookupd_1,nsqlookupd_2,nsqadmin
    ```
    $ forego start -f Procfile.nsqd
    forego        | starting nsqlookupd_2.1 on port 5000
    forego        | starting nsqlookupd_1.1 on port 5100
    forego        | starting nsqd_1.1 on port 5300
    forego        | starting nsqd_2.1 on port 5600
    forego        | starting nsqadmin.1 on port 6000
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.318010 nsqlookupd v0.3.8 (built w/go1.6.2)
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.319185 TCP: listening on [::]:4162
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.319294 HTTP: listening on [::]:4163
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.323656 nsqlookupd v0.3.8 (built w/go1.6.2)
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.324720 TCP: listening on [::]:4160
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.324771 HTTP: listening on [::]:4161
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.334497 nsqd v0.3.8 (built w/go1.6.2)
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.334885 ID: 679
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.335107 TOPIC(test): created
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.335451 TOPIC(test): new channel(nsq_to_file)
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.335475 NSQ: persisting topic/channel metadata to nsqd-2-data/nsqd.679.dat
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.336895 TCP: listening on [::]:4152
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.336925 HTTP: listening on [::]:4153
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.337853 LOOKUP(0.0.0.0:4160): adding peer
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.337863 LOOKUP connecting to 0.0.0.0:4160
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.336840 nsqd v0.3.8 (built w/go1.6.2)
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.337036 ID: 678
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.337253 TOPIC(test): created
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.337534 TOPIC(test): new channel(nsq_to_file)
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.337556 NSQ: persisting topic/channel metadata to nsqd-1-data/nsqd.678.dat
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.338523 HTTP: listening on [::]:4151
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.338578 TCP: listening on [::]:4150
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.338872 TCP: new client(127.0.0.1:61984)
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.338898 CLIENT(127.0.0.1:61984): desired protocol magic '  V1'
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.339109 CLIENT(127.0.0.1:61984): IDENTIFY Address:antoniko-3.local TCP:4152 HTTP:4153 Version:0.3.8
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.339128 DB: client(127.0.0.1:61984) REGISTER category:client key: subkey:
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.339603 LOOKUP(0.0.0.0:4160): adding peer
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.339616 LOOKUP connecting to 0.0.0.0:4160
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.340208 TCP: new client(127.0.0.1:61985)
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.340225 CLIENT(127.0.0.1:61985): desired protocol magic '  V1'
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.340302 CLIENT(127.0.0.1:61985): IDENTIFY Address:antoniko-3.local TCP:4150 HTTP:4151 Version:0.3.8
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.340313 DB: client(127.0.0.1:61985) REGISTER category:client key: subkey:
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.340709 LOOKUPD(0.0.0.0:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:0.3.8 BroadcastAddress:}
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.340747 LOOKUP(0.0.0.0:4162): adding peer
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.340751 LOOKUP connecting to 0.0.0.0:4162
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.340431 TCP: new client(127.0.0.1:61986)
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.341293 DB: client(127.0.0.1:61984) REGISTER category:topic key:test subkey:
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.341547 LOOKUPD(0.0.0.0:4162): peer info {TCPPort:4162 HTTPPort:4163 Version:0.3.8 BroadcastAddress:}
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.341585 LOOKUPD(0.0.0.0:4160): topic REGISTER test
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.340450 CLIENT(127.0.0.1:61986): desired protocol magic '  V1'
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.340663 CLIENT(127.0.0.1:61986): IDENTIFY Address:antoniko-3.local TCP:4152 HTTP:4153 Version:0.3.8
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.340694 DB: client(127.0.0.1:61986) REGISTER category:client key: subkey:
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.341183 TCP: new client(127.0.0.1:61987)
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.341208 CLIENT(127.0.0.1:61987): desired protocol magic '  V1'
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.341376 CLIENT(127.0.0.1:61987): IDENTIFY Address:antoniko-3.local TCP:4150 HTTP:4151 Version:0.3.8
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.341391 DB: client(127.0.0.1:61987) REGISTER category:client key: subkey:
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.341548 DB: client(127.0.0.1:61986) REGISTER category:topic key:test subkey:
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.342071 LOOKUPD(0.0.0.0:4162): topic REGISTER test
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.342103 DB: client(127.0.0.1:61986) REGISTER category:channel key:test subkey:nsq_to_file
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.342195 DB: client(127.0.0.1:61987) REGISTER category:topic key:test subkey:
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.339669 LOOKUPD(0.0.0.0:4160): peer info {TCPPort:4160 HTTPPort:4161 Version:0.3.8 BroadcastAddress:}
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.339683 LOOKUP(0.0.0.0:4162): adding peer
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.339686 LOOKUP connecting to 0.0.0.0:4162
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.340982 LOOKUPD(0.0.0.0:4162): peer info {TCPPort:4162 HTTPPort:4163 Version:0.3.8 BroadcastAddress:}
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.341028 LOOKUPD(0.0.0.0:4160): topic REGISTER test
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.341402 LOOKUPD(0.0.0.0:4162): topic REGISTER test
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.341750 LOOKUPD(0.0.0.0:4160): channel REGISTER test nsq_to_file
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.341951 LOOKUPD(0.0.0.0:4162): channel REGISTER test nsq_to_file
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.342277 LOOKUPD(0.0.0.0:4160): REGISTER test nsq_to_file
    nsqd_2.1      | [nsqd] 2016/11/19 23:08:38.342551 LOOKUPD(0.0.0.0:4162): REGISTER test nsq_to_file
    nsqlookupd_2.1 | [nsqlookupd] 2016/11/19 23:08:38.342687 DB: client(127.0.0.1:61987) REGISTER category:channel key:test subkey:nsq_to_file
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.341874 DB: client(127.0.0.1:61984) REGISTER category:channel key:test subkey:nsq_to_file
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.341916 DB: client(127.0.0.1:61985) REGISTER category:topic key:test subkey:
    nsqlookupd_1.1 | [nsqlookupd] 2016/11/19 23:08:38.342432 DB: client(127.0.0.1:61985) REGISTER category:channel key:test subkey:nsq_to_file
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.342289 LOOKUPD(0.0.0.0:4160): channel REGISTER test nsq_to_file
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.342515 LOOKUPD(0.0.0.0:4162): channel REGISTER test nsq_to_file
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.342791 LOOKUPD(0.0.0.0:4160): REGISTER test nsq_to_file
    nsqd_1.1      | [nsqd] 2016/11/19 23:08:38.342945 LOOKUPD(0.0.0.0:4162): REGISTER test nsq_to_file
    nsqadmin.1    | [nsqadmin] 2016/11/19 23:08:38.345705 nsqadmin v0.3.8 (built w/go1.6.2)
    nsqadmin.1    | [nsqadmin] 2016/11/19 23:08:38.346025 HTTP: listening on [::]:4171
    ```
# Тестируем на получение дублей сообщений

*  `./publish.sh` - пишем в nsqd_1 и в nsqd_2 по 1000 сообщений
   ```
   $ ./publish.sh
   OKOKOKOKOKOKOKOKOKOKOKOKOKOKOKOKOKOKOKOKO...
   ...
   OKOKOKOK....
   ```
*  ` forego start -f Procfile.nsqtofile ` - запускаем nsq_to_file_1,nsq_to_file_2 (получаем наши сообщения) и запоминаем как называются файлы с логами
   ```
   $ forego start -f Procfile.nsqtofile
   forego  | starting file_1.1 on port 5000
   forego  | starting file_2.1 on port 5100
   file_2.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=test
   file_1.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=test
   file_1.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] (antoniko-3.local:4152) connecting to nsqd
   file_2.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] (antoniko-3.local:4152) connecting to nsqd
   file_1.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] (antoniko-3.local:4150) connecting to nsqd
   file_2.1 | 2016/11/19 23:04:54 INF    1 [test/nsq_to_file] (antoniko-3.local:4150) connecting to nsqd
   file_1.1 | 2016/11/19 23:04:54 INFO: opening nsq_to_file_1/test.antoniko-3.2016-11-19_23.log
   file_1.1 | 2016/11/19 23:04:54 syncing 1 records to disk
   file_2.1 | 2016/11/19 23:04:54 INFO: opening nsq_to_file_2/test.antoniko-3.2016-11-19_23.log
   file_2.1 | 2016/11/19 23:04:54 syncing 1 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 136 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 145 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 164 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 165 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 162 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 164 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 105 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 135 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 164 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 164 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 134 records to disk
   file_2.1 | 2016/11/19 23:04:54 syncing 150 records to disk
   file_1.1 | 2016/11/19 23:04:54 syncing 104 records to disk
   file_2.1 | 2016/11/19 23:05:24 syncing 102 records to disk
   file_1.1 | 2016/11/19 23:05:24 syncing 4 records to disk
   file_2.1 | 2016/11/19 23:06:02 INF    1 [test/nsq_to_file] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=test
   file_1.1 | 2016/11/19 23:06:10 INF    1 [test/nsq_to_file] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=test
   ```
* `tests-nsqd  -a=nsq_to_file_1/test.antoniko-3.2016-11-19_23.log -b=nsq_to_file_2/test.antoniko-3.2016-11-19_23.log` - сравниваем логи на повторные сообщения
   ```
    $ tests-nsqd  -a=nsq_to_file_1/test.antoniko-3.2016-11-19_23.log -b=nsq_to_file_2/test.antoniko-3.2016-11-19_23.log
    message |count
    _____________________
    _____________________
    total uniq in A: nsq_to_file_1/test.antoniko-3.2016-11-19_23.log 986
    total uniq in B: nsq_to_file_2/test.antoniko-3.2016-11-19_23.log 1014
    total doubles -  0
    total uniq -  2000
   ```

Можно повторять эти три шага - дублей нет