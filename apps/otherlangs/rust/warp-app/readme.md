




same / json (warp) 17000
http://0.0.0.0:3200/json_warp
http://0.0.0.0:3200/json

big-text 5kb 42000-44500
http://0.0.0.0:3200/test


small-json (warp) - 55000
http://0.0.0.0:3200/smalljson_warp

json-text 55000
http://0.0.0.0:3200/json_text

small-text 50000-60000
http://0.0.0.0:3200/

404 - 63000
http://0.0.0.0:3200/sdfsaffd


1) ai - models
2) lang - jai, odin, nim
3) framework - finata scala, nim test, jai, odin
4) json - elixir plug, akka/finata scala, .net core, haskell, ruby, (crystal - bad)



место
rust, java, haskell, scala, clojure, dotnet, node, ruby, php, python, go, exlixir, crystal, lua

установка запуск
node (pm2), go, crystal, dotnet, clojure, elixir, zig/c, rust (docker 4+ min), java, scala, php (docker 4+ min), flutter + dart, ruby (10 min +), haskell (5-10 min+), C++, kotlin native, unity, unreal engine

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

go   30 sec (52 sec)
rust 223 sec (30-40 sec carets) (264 sec)

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