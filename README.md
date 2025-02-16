#### test frameworks 2021


rust actix/warp download/build                    80-100k+ (210k)   2 min ~   (docker bad)
polkadot js                                       30-87k+  (156k)   0 min ~   2-4 repl / pm2
go fiber download/build         (j 47k)           52-72k   (170k)   5-30 sec ~ (docker bad 40-50k)
php8.1 workerman / mark / webman (j 55k-70k)      46-51k   (88k)    3-4 workers (proccess, docker ok)
HAproxy                                           40-64k   (140k)
java activej                         (j 26k)      35-64k   (130k)               (docker bad 33k)
sanic py v3.11                 (j 30k) 14(21) -  (35)57k   (60k)    2sec - 1min~ 3 repl / docker
nginx                                             29-51k   (60k)
php8.1 simps / fomo / swoole                         40k   (82k)    ...        (docker ok, swarm bad) без потерь
scala akka download/build                            37k   (85k)    20-40 sec  (docker bad)
clojure donkey                (j 20k)     30k   (62-75k)  (172k)       0 min      (docker ok, swarm bad) без потерь
crystal routes.cr, toro download/build    29k      (33k)    35 sec+    (docker bad, swarm bad)
express js                                7-24k    (24k)    6-9 repl / pm2
bun js                                    20k
haskell warp/wei download/build           20k+     (24k)    5+ min (docker bad)
c# .net core - fastendpoint       (j 12k) 12k      (14k)

vala/d lang - like gcc
nim/zig - don't work

1) rust warp/actix  text                  json                         404 180000       build 2 min+
2) php8.1           text 63000            json 55000 - 70000 (59000)   404 130000       create docker - 2-3 min
3) java activej     text 45000            json 55000 -                 404 112000       load packages 30-40 sec~, build 10 s
4) go fiber         text 35000            json 47000+                  404 81000        build 1-5 sec
5) js-node polkadot text 47000            json 37000(9500  text)       404 41000-81000  build 0 sec
6) scala akka       text 37000            json                         404 85000        build 8-15 sec
7) py sanic         text 35000(14000 x1)  json 30000(11200 x1)         404 24000(12000) create docker - 20 sec - 1 min~
8) clojure          text 30000            json                         404 62000-72000  run 0 sec
9) crystal routes   text 29000            json                         404 33000        run 5-8 sec, build 37 sec
10) haskell         text 20000            json                         404 24000        run 3 sec, first build 5 min+
11) C# FastEndpoints text 11500           json 11500(14000)                  404 7000         run 0 sec, first build 3 sec


json
polkadot =      9500  - 37000 (json) -         26000 - 47000( text)        404 41000 - 81000     +20%~ +30%
1 repl = json   15500         (big 9500)  text 26000 - 30000               404 41000
2-3 repl = json 30000 - 37000 (big 26333) text 38000 - 47000+              404 81000
sanic =         11200 - 30000 (json) -         14000 - 35000 (text)        404 11000 - 24000
1 repl = json   11300 text                     14000 - 21000               404 11000
3 repl = json   21000 - 30000 text             28000 - 35000               404 24000
go fiber =      38000 - 47000 (fiber/json) -   29000 - 35000 (text/json)   404 81000
php             52000 - 70000 (json)           63000         (text)        404 130000


Compile time:

1) go       = 0 sec < 15 sec              (fast build)
2) clojure  = 0 sec - 2 sec
3) nim, zig = 0 sec ~                     (load, build) don't work
4) C#, .Net Core 7 - mvc - 1-3 sec        load, fast build
3) python   = 0 sec - 15 sec              load
4) node.js  = 0 sec < 1 min               load
5) crystal  = 5 sec / 35+ sec             (build)
6) scala    = 10 sec < frist 40 sec~      (download & build)
7) php      = 0 sec, fisrt 3 min - 5 min+ (download & build)
8) java     = 10-20 sec                   (load, build)
9) rust     = 2 min +                     (download & build)
10) haskell  = 5 min +                    (download & build)
11) c++
12) llvm

Easy code;

1) python
2) js / node
3) php
4) clojure
5) .Net Core 7 - mvc
6) Scala
7) go
8) crystal
9) rust

Nice coding:

1) crystal
2) Go
3) .Net Core 7
4) Clojure
5) python
6) js
7) rust
8) scala

Speed code:

1) crystal
2) python
3) js
4) php
5) clojure
6) go
7) scala
nim ?
8) java
9) rust
zig / d / c ?
10) haskell
11) C++
12) llvm


## service starting, port for each application 0.0.0.0 и 3200


### Sanic - 12500 - 15275.80 rps
```

docker-compose -f api_sanic.yml up

```


### aiohttp - 4812.64 rps
```

docker-compose -f api_aiohttp.yml up

```

### fastapi - 3800.24 rps
```

docker-compose -f api_fastapi.yml up

```


## testing RPS:

```
curl http://0.0.0.0:3200/test?test=test

docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
```

(без принтов)

# table 1

kafka =                                           1 mln (+ replics) x 500kb        (gRPC)
📌 clickhouse batch insert                        500k query/s, DB, logs (x60 vs postgres)
grafana loki ?                                            ~
📌 polkadot js  (2 repl * pm2)                (30)47 - 66k (text/json)
📌 rust (actix, warp) (1) =                       52 - 64k, (80k, 100k) API не реплецируется в swarm
📌 go (fiber) (1) =                               16k, 56k (16-51k in prod) API не реплецируется в swarm
redis = 8000 (py) - 30000 (go) -                  40k - 50k (go, py) batch (RAM, DB)
php8.1 workerman / mark                           46k - 50k (88k)
📌 sanic py 3.11, repl3 * swarm =                 10k (17k) - 21k (45k) API
nginx =                                           18k - 60k (29-37k) (RAM, cache)
HAproxy   (docker)                                40k~
php8.1 simps / fomo                               40k (82k)
clojure donkey                                    30k  (62-72k) 0 min не реплицируется в swarm
rabbit =                                          30k - 1-3 mln (batch 100-1000 x 1 message) (gRPC)
📌 aiohttp py repl3 * swarm                       4.8k (9k) - 28k API
traefik    (rust,go)                              27k  40-57%
exoress js (6-9)                                  24K API
lowhttpsever js    / bun js                       20k API не реплецируется
haskell warp/wei                                  20k API не реплецируется
postgres batch insert * A/B async                 6 k - 9.4k  query/s (x2) # >500 batch прерывание
postgres batch insert                             2.5 k - 8k  query/s (x2)
GET postgres pool select 1 row                    3.4 - 4 req/s API, query/s
fastapi py repl3                                  3.8 (8.3) (500-1400 in prod) API
node express                                      ~
(api) sanic insert 1 row pool                     900-1300 query/s (x2) API, benchmark A/B, simple



polkadot in docker -  30000
actix in docker -     33383
fiber in docker -     30000
nginx (1-2)       -     15902-16798.91

# table 2

min - max (5 kb text)
polkadot.js (text/json) 30000 (1) -     66153(47160)            (2-4) + https://github.com/lukeed/polkadot
warp (rust)(1) text     55104   (122575)63540       83-96-134%  (1) не реплицируется
actix (rust) (1) text   52000.75        60200       78-90-126%  (1) не реплицируется
fiber (go)(1) text      40000   (107000)51605.92    60-78-109%  (1) не реплицируется
php8.1 workerman / mark 46000  - (88000)50000                       3-4 worker (proccess)
sanic -                 17367 (1)   -   45000       26-68-95.4% (3) +
HAproxy         (rust,go)  (110000-86000)40746       61-86%
php8.1  simps / fomo              (82000)40500       61-80%         не реплицируется в swarm
 rocket (rust)(1)       30273.00  (88437)36841       45-55-78%   (1) text (native), не реплицируется
 nginx (rust,go) (1)               (37000)29000      43-55.9-78% 
 aiohttp                 9717.09 (1) -   28609.18    14-43-60%   (3) +
 traefik                       (rust,go) 27000          40-57%
 express.js              7200 (1)    -   24000       10-36-51%   (6-9) запускать через pm2 start express.js -i 6
 lowhttpserver           pm2, без реплик-20600       31-43%      (1) запускать только через pm2,  не реплицируется
 hyperexpress            pm2, без реплик-20000       31-43%      (1) запускать только через pm2,  не реплицируется
 bun.js                  pm2, без реплик-20000       31-43%      (1) запускать только через pm2,  не реплицируется
  fastrouter (go) (1)              (94831)16800.12    23%        (1) не реплицируется
  rayo.js                 11912.70 (1)    15377.73    18-23%     (1) не реплицируется
  echo (go)(1) text                (98912)14703       22%        (1) не реплицируется
  blacksheep              5622.84 (1) -   8710.87     8-13%      (2) не реплицируется (больше 8000)
  fastapi                 4290.38 (1) -   8371.09     6.4-12% (2) не реплицируется (больше 8000)


без реплик
actix (rust) (1) text/javascript    52000.75        59635 не реплицируется
fiber (go)(1) text                                  52605.92 не реплицируется
php8.1 workerman / mark46000  - (88000)50000                       3-4 worker (proccess)
warp (rust)(1) text (native) DEV             (63852)47071.30 не реплицируется
(vad) actix (rust)(1) text (docker)  41000           44866.24 не реплицируется
php8.1  simps / fomo              (82000)40500       61-80%         не реплицируется в swarm
(bad) fiber (go)(1) text (docker)    40000           44000 не реплицируется
fiber (go)(1) json                  39077           40000 не реплицируется
rocket (rust)(1) text (native)      30273.00        (88437)36841 не реплицируется
rocket (rust)(1) text (docker)      30273.00        (69529)35841 не реплицируется
polkadot.js                         30000 -         30000 https://github.com/lukeed/polkadot
rocket (rust)(1) text (swarm)       25768           (57663.22)25768 не реплицируется
hyperexpress.js (pm2/node)          18500-          20600  не реплицируется
lowhttpserver.js (pm2/node)         18724.52 -      20000  не реплицируется
bun.js                              19252.44        20000  не реплицируется
haskel scotty                       19000           19000  не реплицируется
sanic -                             9000-           17367.74
restana.js (1) (pm2/node)           13309.37        16767.81 (bad) не реплицируется
rayo.js (1) (pm2/clustr)            11912.70        15377.73 (ok) не реплицируется
fastrouter (go) (1)                                 15084.12
echo (go)(1) text                            (98912)14703 не реплицируется
(bad -) bun.js (docker)              13116.88 -      14254.47 не реплицируется
(bad -) polkadot.js (docker)         12599.14 -      13563.45 не реплицируется
swarm (3) express.js                 12765.74 -      13558.51 не реплицируется
(bad -) hyperexpress.js (docker)                    12511.53 не реплицируется
(bad -) lowhttpserver.js(docker)1    2945.04 -       13641.80 не реплицируется
aiohttp -                           5131 -          9717.09
express.js                          6200            7200
blacksheep -                        4506.06 -       5622.84 не реплицируется (больше 8000)
fastapi -                           2990 (3300) -   4290.38 не реплицируется (больше 8000)

### with replication
📌 pm2/cluster (2-3) polkadot.js55947    66153.29(46371.22) (good) text / json
pm2/cluster (3) polkadot.js     52000      66153.29
pm2/cluster (8) polkadot.js     45311.84 - 50790.42
📌 fiber (go)(1) text                      51605.92 (no docker)
php8.1 workerman / mark /webman 46000  - (88000)50000                       3-4 worker (proccess)
swarm (3) sanic -            35084.23 -  44964.92 (good) x.1.5 (aio repl), x2.5 - x4.5 (sanic),  x5.7 (fastapi repl), x10.4 (fastapi)
php8.1  simps / fomo / swoole  (82000)40500       61-80%         не реплицируется в swarm
swarm (5) sanic -               33905.56 -  41749.31 (good) 
(bad -) swarm (5) polkadot.js   33000.00 -  43681 (bad) , больше 5 не эффективно и жрет ресурсы
(bad -) swarm (4) polkadot.js   33000.00 -  40580 (bad) 
rocket (rust)(2) text (swarm)   25768       (60685.42)26527.31
swarm (2) sanic -               35931.97 -  37589.83 (ok)
(bad -) swarm (3) polkadot.js   30000.00 -  37946 (bad)  
swarm (4) sanic -               33797.64 -  37016.14 (ok)
rocket (rust)(1) text (native)  30273.00    (88437)36841
rocket (rust)(1) text (docker)  30273.00    (69529)35841
rocket (rust)(1) text (swarm)   25768       (57663.22)25768
swarm (6) sanic -               32696.34 -  35199.79
swarm (7) sanic -               27454.78 -  31079.01
(bad -) swarm polkadot.js(2)                30425.00 (bad)
📌 swarm (3) aiohttp -          17700    -  28609.18 (good) 63% of sanic, x2.9 (aio), x3.4 (fastapi)
swarm (4) aiohttp -             19062.54 -  25170.81 (ok)
swarm (5) aiohttp -             22691.59 -  25116.64
express.js (9)(pm2/clustr)      21000 -     23900
haskell warp/wei                20000
📌 express.js (6)(pm2/clustr)   19893 -     23822.64 - реплицируется в swarm/cluster, грузит проц (ok)
express.js (8)(pm2/clustr)      17926 -     21815.61
📌 express.js (5)(pm2/clustr)   17926 -     20500.61
📌 hyperexpress(1)              18500-      20600 (не реплицируется в докере/pm2 cluster), (ok) (pm2/node)
📌 lowhttpserver(1)             18724.52 -  20000 (не реплицируется в докере/pm2 cluster), (ok) (pm2/node)
📌 bun.js (1)(pm2/node)         18116.88 -  20000 (не реплицируется в докере/pm2 cluster), (ok) (pm2/node)
(bad -) swarm (5) express.js    13293.12 -  18193.73
(bad -) swarm (6) express.js    14999.26 -  18193.73
swarm (6) aiohttp -             17543.09    
📌sanic (1) -                   9000-       17367.74 (ok), (60% of aio repl), x1.78 (aio), x2 (fastapi repl), x4 (fastapi))
express.js (3)(pm2/clustr)      16522.89-   17247
(bad -) restana.js (2) (pm2/node) 15557       16767.81 
swarm (4) express.js            14827.17 -  16876.24
(bad -) rayo.js (2) (pm2/clustr)12405.12    16143.17 (не реплецируется)
rayo.js (1) (pm2/clustr)        11912.70    15377.73 (ok)
swarm (2) aiohttp -             11000 -     15451.03
(bad -) http0.js(2) (pm2/clustr)12599.44 -  15086
(bad -) bun.js (docker)         13116.88 -  14254.47 
.net core 7 FastEndpoints
(bad -) polkadot.js (docker)    12599.14 -  13563.45
swarm (3) express.js            12765.74 -  13558.51
(bad -) hyperexpress.js (docker)12511.53
(bad -) lowhttpserver.js(docker)12945.04 -  13641.80 
c++ (lib) evhtp                 11300
http0.js -  13                  11058.35
aiohttp (1) -                   5131 -      9717.09 (55% of sanic), x1.16 (fastapi repl), x2.2 (fastapi)
(baa -) swarm (3)haskell warping            9200
(bad -) swarm (2)express.js                 9387.48
(bad -) swarm (3) ban.js                    9492.40
(bad -) swarm (2) hyperexpress              9492.40
(bad -) swarm (2) lowhttpserver             9145
swarm (2) blacksheep -          7400-       8710.87-
swarm (2) fastapi -             6900 -      8371.09 (18-20% of sanic repl), (29% of aio repl), (43% of sanic), (86% of aio)
swarm (3) blacksheep -          7672.53 
swarm (3) fastapi -             6133 -      7575.02 
express.js (1) pm2              6403 -      7200
swarm (5) fastapi -                         6807.05
swarm (5) blacksheep -                      6500.53 
blacksheep (1) -                4506.06 -   5622.84 
swarm (6) fastapi -                         5131
(bad -) express.js              12945.04 -  5000  (docker)
fastapi (1) -                   2990 (3300)-4200
clang cerveur                   2474



















45053




## aiohttp (optimal 3)
```
// init
docker swarm init
docker-compose -f swarm_api_aiohttp.yml build
docker stack deploy -c swarm_api_aiohttp.yml swarm_test_aiohttp_api
docker service ls

// change replica
docker service scale swarm_test_aiohttp_api_swarm_test_aiohttp_api=3
docker service ls

// tests
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 400000 http://0.0.0.0:3200/test?test=test

// remove service
docker service rm [tab] 
```

## fastapi (optimal 2)
```
// init
docker swarm init
docker-compose -f swarm_api_fastapi.yml build
docker stack deploy -c swarm_api_fastapi.yml swarm_test_fastapi_api
docker service ls

// change replica
docker service scale swarm_test_aiohttp_api_swarm_test_aiohttp_api=3
docker service ls

// tests
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test


// remove service
docker service rm [tab] 
```


## sanic (optimal 3)
```
// init
docker swarm init
docker-compose -f swarm_api_sanic.yml build
docker stack deploy -c swarm_api_sanic.yml swarm_sanic_api
docker service ls

// change replica
docker service scale swarm_test_aiohttp_api_swarm_test_aiohttp_api=3
docker service ls

// tests
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 400000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost:3200/test?test=test

// remove service
docker service rm [tab] 
```

## blacksheep (optimal 2)
```
// init
docker swarm init
docker-compose -f swarm_api_blacksheep.yml build
docker stack deploy -c swarm_api_blacksheep.yml swarm_test_blacksheep_api
docker service ls

// change replica
docker service scale swarm_test_aiohttp_api_swarm_test_aiohttp_api=3
docker service ls

// tests
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test

// remove service
docker service rm [tab] 
```

## express js (optimal 6,9)
```

// cluster
cd apps/otherlangs/node/natural/
pm2 start express.js -i max
pm2 scale express 6
pm2 restart express.js -i 6

//pm2 start express.js 6

docker run  --network host --rm jordi/ab -k -c 100 -n 300000 http://localhost:3000/?test=test

// rm cluster
pm2 delete express

node cluster.js
https://medium.com/geekculture/scaling-node-js-applicationswith-pm2-clusters-c216c4468d66



// init
docker swarm init
docker-compose -f swarm_api_node_hyperexpress.yml build
docker stack deploy -c swarm_api_node_hyperexpress.yml swarm_test_hyperexpress_api
docker service ls

// change replica
docker service scale swarm_test_hyperexpress_api_swarm_test_hyperexpress_api=5
docker service ls

// tests
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 200000 http://0.0.0.0:3200/test?test=test

// remove service
docker service rm [tab] 
```


## polkadot js (optimal 2)
```

// cluster
cd apps/otherlangs/node/natural/
pm2 start express.js -i max
pm2 scale express 2
pm2 restart polka.js -i 2

//pm2 start express.js 2

docker run  --network host --rm jordi/ab -k -c 100 -n 700000 http://localhost:3000/test/?test=test

// rm cluster
pm2 delete express

```

## HAproxy experiment - 30k * 50k * 60k
```
cd apps/otherlangs/rust/actix-app
cargo build --release
./target/release/superapp

cd apps/otherlangs/go/fiber/superapp
go build -o ./bin/server ./cmd
./bin/server

docker-compose -f haproxy_experiments.yml up -d


// 1
curl http://localhost:3000/test
// 2
curl http://localhost:3200/test
// balancer
curl http://all.localhost/test
// 404
curl http://all.localhost/sfsfe

// rps
# 1 fiber - 49957.73 / 117594.37
docker run --network host --rm jordi/ab -k -c 100 -n 600000 http://localhost:3000/test
docker run --network host --rm jordi/ab -k -c 100 -n 1300000 http://localhost:3000/fasffewf
# 2 actix - 59381.96 / 130528.95
docker run --network host --rm jordi/ab -k -c 100 -n 700000 http://localhost:3200/test
docker run --network host --rm jordi/ab -k -c 100 -n 1400000 http://localhost:3200/fasffewf
# 3 HAproxy balancer 39830.72 / 85283.02
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://all.localhost/test
docker run --network host --rm jordi/ab -k -c 100 -n 1000000 http://all.localhost/errrrrrrrrrrrrrrrr

// stop
docker-compose -f haproxy_experiments.yml down

```


## Nginx experiment - 30k * 50k * 60k
```
cd apps/otherlangs/rust/actix-app
cargo build --release
./target/release/superapp

cd apps/otherlangs/go/fiber/superapp
go build -o ./bin/server ./cmd
./bin/server

docker-compose -f nginx_experiments.yml up -d

// 1
curl http://localhost:3000/test
// 2
curl http://localhost:3200/test
// balancer
curl http://all.localhost/test
// 404
curl http://all.localhost/sfsfe

// rps
# 1 fiber - 49957.73 / 117594.37
docker run --network host --rm jordi/ab -k -c 100 -n 600000 http://localhost:3000/test
docker run --network host --rm jordi/ab -k -c 100 -n 1300000 http://localhost:3000/fasffewf
# 2 actix - 59381.96 / 130528.95
docker run --network host --rm jordi/ab -k -c 100 -n 700000 http://localhost:3200/test
docker run --network host --rm jordi/ab -k -c 100 -n 1400000 http://localhost:3200/fasffewf
# 3 nginx balancer 32241 / 36228.30
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost/test
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost/

// stop
docker-compose -f nginx_experiments.yml down


```

## Traefik experiment - 30k * 50k * 60k
```
cd apps/otherlangs/rust/actix-app
cargo build --release
./target/release/superapp

cd apps/otherlangs/go/fiber/superapp
go build -o ./bin/server ./cmd
./bin/server

docker-compose -f traefik_experiments.yml up -d

// 1
curl http://localhost:3000/test
// 2
curl http://localhost:3200/test
// balancer
curl http://host.localhost/test
// 404
curl http://host.localhost/sfsfe

open in browser http://traefik.localhost/dashboard/#/

// rps
# 1 fiber - 49957.73 / 117594.37
docker run --network host --rm jordi/ab -k -c 100 -n 600000 http://localhost:3000/test
docker run --network host --rm jordi/ab -k -c 100 -n 1300000 http://localhost:3000/fasffewf
# 2 actix - 59381.96 / 130528.95
docker run --network host --rm jordi/ab -k -c 100 -n 700000 http://localhost:3200/test
docker run --network host --rm jordi/ab -k -c 100 -n 1400000 http://localhost:3200/fasffewf
# 3 traefik balancer 27439 / 39686
docker run --network host --rm jordi/ab -k -c 100 -n 400000 http://host.localhost/test
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://host.localhost/sdfsdfd

// stop
docker-compose -f traefik_experiments.yml down


```


<!-- ## nginx 
```

// init
docker swarm init
docker-compose -f swarm_nginx_experiments.yml build
docker stack deploy -c swarm_nginx_experiments.yml swarm_test_nginx_balancer
docker service ls

// change replica
docker service scale [tab]
docker service ls

// tests
curl http://localhost/test
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost/sfgerfewrger
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost/test

// remove service
docker service rm [tab] 
``` -->


## tests
```
docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 400000 http://0.0.0.0:3200/test?test=test
docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost:3200/test?test=test
```

## conclusion
## swarm exit
```
docker swarm leave --force
docker service ls

```


Версию python можно изменить в ```dockers/python/Dockerfile```

