1) ai - models
2) lang - jai, odin, nim
3) framework - finata scala, nim test, jai, odin
4) json - elixir plug, akka/finata scala, .net core, haskell, ruby, (crystal - bad)



A)

получается sanic x6 + orjson = +x2~
плюс версия 3.11 +10-60%

в итоге 3 реплики бьют go по статичной отдаче

Б)

если брать fastapi + стандартный json + версия 3.10, без реплик

то будет 1/10 от го


эффективность
rust, clojure, python, scala, nodejs, elixir, php

целесообразность

1) python   (6k-fa       30k-58k            51k-j      32-58k)     =       51k-j  (3)
 - go       (40-52k      52k(46k)           58k-j   130k(110k),   =       58k-j  (1-8)    =+13%    (x200% без реплик)
 - php      (                55k            61k-j        125k+)   =       61k-k+ (4)      =+19%    (x200% без реплик)
 - rust     (42k-ro      72k(55k)     95k-j(75k-j)  210k(132k)    = 95k-j(75k-j) (1-8)    =+25-86% (x316% без реплик)
2) 
 - node     (20-32k-ex   72k(47k)     67k-j(48k-j)  162k(102k)    = 67k-j(48k-j)  (3-7)   =-6% -/x2.5 (+60% - x220% без реплик)
 - exixir,                   60k              x-j        100k     =         x-j           =+0%  (slower in docker)
 - clojure                   52k            35k-j        120k     =        35-j           =-45% (slower in docker)
3) scala                     44k+             x-j        103k     =         x-j           =X    (slower in docker)
 - java,                     75k            48k-j        189k     =       48k-j           =-6%  (slower in docker)
 - zig/c/c++
4) crystal                   53k              x-j         59k     =         x-j           =X    (slower in docker)
 - nim
 - jai, odin, lua 


№ язык          запуск         красота кода    разработка  скорость  функциональность  вес           востребованность/оплата

python ++       средне/долго   хорошо          оч. легко   быстро    очень             средний       очень/высокая
nodejs +++      быстро         средне          оч. легко   средне    очень             легковесно    высокая/средняя
elixir          быстро/средне  красивый        ?ок         быстро    +?ок              легковесно    ?ок/?
clojure +-      быстро         ок              легко,      средне    +норм             легковесно    ?ок/  очень
rust ++         быстро/долго   красивый        сложно      быстро    очень             высокий       очень/очень
scala           быстро         средне          средне      средне    выскоая           много         высокая/очень
go ++           быстро         средне          сложно      быстро    высокая           средний       высокая/высокая
php             быстро/долго   некрасивый      легко       быстро    web               легковесно    средняя/низкая
crystal         быстро/средне  красивый        ?неок       средне    -н изко            легковесно    -низко/?
c#.netcore +-   быстро         красивый        средне      медленно  +высокая          средний       средне/средняя
java            долго/долго    длинный         сложно      средне    очень             средне        высокая/высокая
haskell         быстро/долго   некрасивый      очень       медленно  высокая           высокий       низкая/очень
zig/c/c++       долго/долго    очень плохо     очень       очень     гипер             высокий       среднее/среднее
ruby            быстро/долго   красивый        ?неок       медленно  высокая           легковесно    -средне/средняя

№ язык          запуск         красота кода      разработка  скорость  вес

elixir          быстро/средне  красивый          ?ок         быстро    легковесно
crystal         быстро/средне  красивый          ?неок       средне    легковесно
ruby            быстро/долго   красивый          ?неок       медленно  легковесно

1 clojure +-    быстро         ок                легко,      средне    легковесно
2 nodejs +++    быстро         средне            легко,      средне    легковесно
3 scala         быстро         средне            средне      средне    много
4 go ++         быстро         средне            сложно      быстро    средний
5 python ++     средне/долго   хорошо            легко       быстро    средний
6 java          долго/долго    длинный           сложно      средне    средне
7 php           быстро/долго   некрасивый        легко       быстро    легковесно
8 c#.netcore +- быстро         красивый          средне      медленно  средний
9 rust ++       быстро/долго   красивый          сложно      быстро    высокий
10 haskell      быстро/долго   некрасивый        очень       медленно  высокий

size 
haskell rust go python scala c#.netcore java php clojure nodejs

(haskell) 240s  (335s)      3.73 GB swarm_haskell_api
(rust)    223s  (250s-264s) 2.32 GB test_frameworks-balancer_rocket_api
(go)      30s   (69s)       1.17 GB test_frameworks-test_fiber_api
(py)      5-85s (60s-130s)  1.17 GB test_frameworks-test_sanic_api
(scala)   3-40s (97s)       1.13 GB swarm_akka_api
(.net)    0s    (351s)      0.87 GB swarm_fastendpoints_api
(java)    31s   (102s)      0.83 GB swarm_activej_api
(php)     0-3s  (267s)      0.76 GB swarm_simps_api
(clojure) 20s   (68s)       0.69 GB swarm_donkey_api
(nodejs)      47s   (70s)       0.32 GB polkadot_api

speed
haskell c#.netcore php rust python java scala nodejs go clojure

clojure 20s (68s) 350+ mb (48 s) load modules (20s) run (0-3s)
go   30 sec                    (69 sec     < 1m)      130 mb+ (39s)
node 47s                       (70s)                  50mb (10s) modules (47s) pm2 (10s)
scala 3-40s (97s)  530mb+ (56s) sbt update (35s~) run (3-6s)
java 31 se (102s) 452mb (70s) mvn packages (31s)
py   5-85~ sec                 (60-130 sec = 2+m)     350 mb+ (60s)
rust 223 sec (30-40 sec carets)(249-264 sec    = 4.5+m)   160 mb+ (26s) upd crates (152s) compile (62s) run (5s)
php  3s (267s)  140mb+ (30s) swoole build (157s) apt install (44s) exts, composer (22s) run (0-3s)
netcore 0s (351s)  280+ mb (330s) netcore (18s) 
haskell 240s (335 s) 625mb (76s) cabal update (23s) build downl (23s) build compile (210s)

место
haskell, rust, go, python, scala, java, dotnet, ruby, php, clojure, node, elixir, crystal, lua

установка запуск
node (pm2), go, crystal, dotnet, clojure, elixir, zig/c, rust (docker 5~ min), java, scala, php (docker 4+ min), flutter + dart, ruby (10 min +), haskell (5-10 min+), C++, kotlin native, unity, unreal engine

запуск
go, crystal, node (pm2), elixir, php, dotnet, ruby, rust, clojure, haskell, java, scala, flutter, unity, unreal engine

простой код
node, crystal, elixir, ruby, php, python, go, dotnet, sacla, java, rust, haskell

json
actix serde, polkadot (3-7), warp, workerman (4), fiber, elixir (?), activej jackson, rocket, fastrouter, dotnet, ruby

text
activej, warp, actix, polkadot (2-7), elixir plug, workerman (4), fiber, crystal, rocket, fastrouter, dotnet, ruby

404
activej, fastendpoints, polkadot (3+), actix, warp, donkey, workerman (4), fiber, fastrouter, rocket, akka, elixir, crystal, express, ruby, sanic, aiohttp, haskell wei/warp, blacksheep, fastapi




404 / max

176702 activej compile       30025.59 [Kbytes/sec] received
155069 c# fastendpoints      18626.51 [Kbytes/sec] received
143000 polkadot (3+) pm2     21029.25 [Kbytes/sec] received
137000 actix release         14761.42 [Kbytes/sec] received
137000 warp release
131985 clojure               8893.53 [Kbytes/sec] received
125197 workerman (4) docker  16138.68 [Kbytes/sec] received
123668 fiber                 19323.22 [Kbytes/sec] received
117474 fastrouter            22944.24 [Kbytes/sec] received
103106 akka                  21749.02 [Kbytes/sec] received
98000+ elixir plug              21053.57 [Kbytes/sec] received (good)
76598 polkadot (6 in docker)    11220.47 [Kbytes/sec] received (slower x2 than w/o docker)
59000 crystal routez         4539.92 [Kbytes/sec] received
404 56570 warp
404 51000 actix
404 49535 rocket release     1395.43 [Kbytes/sec] received
39039-50000 express (4) pm2  20000.55 [Kbytes/sec] received
38105 ruby agoo              2381.98 [Kbytes/sec] received
32000 sanic (3) docker       5725.71 [Kbytes/sec] received
15000 aiohttp (3) docker     2884.75 [Kbytes/sec] received
5731 warp/wei                882.06 [Kbytes/sec] received
http://0.0.0.0:3200/sdfsdf

index small
activej, polkadot.js (3+), C# fastendpoints, actix/warp, workerman, donkey, rocket, akka, elixir, haskell warp/wei, sanic, aiohttp, fastapi, ruby

176702 activej compile          30025.59 [Kbytes/sec] received
168000 polkadot (3+) pm2        24603    [Kbytes/sec] received (polkadot быстрее)
155069 c# fastendpoints         18626.51 [Kbytes/sec] received
147000 actix release            20125.82 [Kbytes/sec] received
140000 warp release             14761.42 [Kbytes/sec] received
132011 actix (in docker)        18564.06 [Kbytes/sec] received (slower 11% than w/o docker)
121148 workerman (4) docker     15616.87 [Kbytes/sec] received
117947 fiber                    18429.22 [Kbytes/sec] received
118622 donkey                   8108.98 [Kbytes/sec] received
103783 rocket release           28885.04 [Kbytes/sec] received
103106 akka                     21749.02 [Kbytes/sec] received
102000 polkadot (6 in docker)   15038.47 [Kbytes/sec] received (slower x2 than w/o docker)
98000 elixir plug               21053.57 [Kbytes/sec] received (good)
84927 virt actix (in docker)    11998.56 [Kbytes/sec] received (slower x1.7)
76790 virt fiber (in docker)    11998.56 [Kbytes/sec] received (slower x1.7)
59000 crystal routez            4539.92 [Kbytes/sec] received
55691 warp
43000 actix
37000 warp/wei haskell          4435.58 [Kbytes/sec] received
32000 sanic (3) docker          5725.71 [Kbytes/sec] received
26000 virt sanic (3) docker     4666.16 [Kbytes/sec] received (slower 60%)
16500 blacksheep (2) docker     2509.98 [Kbytes/sec] received
15000 aiohttp (3) docker        2884.75 [Kbytes/sec] received
12000 fastapi (2) docker        1739.89 [Kbytes/sec] received
http://0.0.0.0:3200/     


text 5kb
java activej, rust warp, node polkadot.js (pm2 7), exlixir plug, sanic, workerman, fiber, crystal, donkey, akka, rocket, express, aiohttp, ruby

74278 activej compile           1458512.34 [Kbytes/sec] received (самый быстрый java)
72000-70000 warp release        1355458.43 [Kbytes/sec] received (warp быстрее)
70000-67000 actix release       1336055.62 [Kbytes/sec] received
72000-54000 polkadot (2-7) pm2  1172055.70 [Kbytes/sec] received (уменьшается при кол-ве реплик)
59812 elixir plug               1180603.44 [Kbytes/sec] received +
58000 sanic (3) docker          1139326.70 [Kbytes/sec] received
55001 workerman (4) docker      1080914.58 [Kbytes/sec] received
54573 actix (in docker)         1073567.72 [Kbytes/sec] received (slower 27% than w/o docker)
52811 warp
54000 fiber                     1061284.56 [Kbytes/sec] received
52682 crystal                   1032144.07 [Kbytes/sec] received
52000 donkey                        1010591.57 [Kbytes/sec] received
47046 polkadot (pm2 7, in docker)   926413.36 [Kbytes/sec] received (slower 48% than w/o docker)
46000 fiber (in docker)             895711.54 [Kbytes/sec] received (slower 20% than w/o docker)
44000 akka                          865728.30 [Kbytes/sec] received
43000 actix                     845142.64 [Kbytes/sec] received
42509 rocket release            838270.23 [Kbytes/sec] received
38480 virt actix (in docker)    726341.62 [Kbytes/sec] received (slower x1.8)
32645 virt fiber (in docker)    642193.26 [Kbytes/sec] received (slower x1.7)
26206-32000 exppress (4-7) pm2  517798.76 [Kbytes/sec] received
29000 aiohttp (3) docker        562732.77 [Kbytes/sec] received
27000 virt sanic (3) docker     525616.48 [Kbytes/sec] received (slower x2.1)
21340 fastrouter                419148.53 [Kbytes/sec] received
17000 c# fastendpoints (docker) 334836.64 [Kbytes/sec] received (практически самый быстрый среди c#, но не самый)
16000 warp/wei                  309487.54 [Kbytes/sec] received
14139 blacksheep (2) docker     278185.98 [Kbytes/sec] received
10500 fastapi (2) docker        203128.77 [Kbytes/sec] received
http://0.0.0.0:3200/test

json
rust actix, polkadot.js (pm2 7), workerman, fiber, sanic orjson/nujson, activej jackson, elixir (?), donkey cheshire, ruby

95000-90000 actix release     579014.17 [Kbytes/sec] received (actix быстрее)
84000 warp release            506297.21 [Kbytes/sec] received
74418 actix (in docker)       456464.39 [Kbytes/sec] received (slower 26% than w/o docker)
55000-67000 polkadot (3-7)    423610.58 [Kbytes/sec] received (увеличивается при кол-ве реплик)
61000 workerman               371875.46 [Kbytes/sec] received 
58000 fiber                   349440.60 [Kbytes/sec] received (same in docker)
51245 sanic orjson (3) docker 316729.48 [Kbytes/sec] received json-text()
50700 virt actix (in docker)  310168.52 [Kbytes/sec] received (slower x3)
48000 activej jackson         290925.34 [Kbytes/sec] received
48000 polkadot (pm2 7, in docker)   297975.10 [Kbytes/sec] received          
44224 sanic nujson (3) docker       273338.30 [Kbytes/sec] received
44000 sanic-json (3) docker         268784.95 [Kbytes/sec] received
43000 sanic ujson (3) docker  264693.60 [Kbytes/sec] received
40500 virt fiber (in docker)   248722.08 [Kbytes/sec] received (slower 43%)
36396 sanic json (3) docker   228685.91 [Kbytes/sec] received
36000 virt sanic (3) orjson   222208.00 [Kbytes/sec] received (slower 42%)
34533 donkey cheshire         209190.79 [Kbytes/sec] received
(finatra don't work, but akka use jackson)
15000 activej json            92356.58 [Kbytes/sec] received
15000 actix                   89488.51 [Kbytes/sec] received
14815 warp
http://0.0.0.0:3200/json