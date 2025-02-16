

0 kafka broker              1m    -       (10-100m batch)
0 clickhouse / loki / log   500k+ / sec   (x 60 postgres)
  1 warp, actix (rust)  min 52000 - 99200 (max 210200) (1 no docker)    max 1
  2 polkadot    (js)    min 30000 - 87615 (max 156074) (2-3 repl * pm2) max 3
  3 fiber       (go)    min 40000 - 72000 (max 162951) (1 no docker)    max 1
  4 sanic       (py)    min 17258 - 57494 (max x)      (3 repl * swarm) max 3+
  5 HAproxy b   -       min 40746 - 63860 (max 147627) (1 * docker)     -
  6 rocket      (rust)  min 30273 - 61000 (max 135000) (1 no docker)    max 1
  7 nginx c/b   -       min 29000 - 50500 (max 61917)  (1 * docker)     -
  8 redis/rabbit c/b    min 30000 - 30000 (max 50000)  (1 * docker)     (1-10 m batch)
  9 traefik b   -       min 27000 - 39837 (max 58273)  (1 * docker)     -
  10 aiohttp     (py)   min 10108 - 30000 (max 30000)  (3 repl * swarm) max 3
    13 express    (js)    min 7200+ - 24000~(max x)      (3-6-9 repl * pm2)   max x
    11 fastrouter (go)    min 16800 - 25515 (max 153713) (1 no docker)        max 1
    12 echo       (go)    min 14703 - 24840 (max 160485) (1 no docker)        max 1
    14 lowhttpserver.js   min 20600 - 20600 ~            (1) не реплицируется - uWS
    15 hyperexpress.js    min 20000 - 20000 ~            (1) не реплицируется - uWS
    16 bun   (js, zig)    min 20000 - 20000 ~            (1) не реплицируется
      17 rayo      (js)         min 11912 - 18386 ~            (1) не реплицируется
      18 postgres batch insert  min 6000  - 20000~             (i in docker)      -
      19 blacksheep(py)         min 5622  - 12000 ~            (2 repl * swarm)   max 2
      20 fastapi   (py)         min 3800  - 11620              (2 repl * swarm)   max 2
        21 postgres pool select 1 row 3400 - 4000+              (1)
        22 postgres pool insert 1 row  900 - 2600~              (1)


min - max (5 kb text)
warp (rust)(1)t   (122575)55104(88000)        63540    (99199)   83-96-134%  (1) не реплицируется (189000/98000)
actix (rust) (1) text     52000(88000)        60200    (95381)   78-90-126%  (1) не реплицируется (210279)
polkadot.js (text/json(1) 30000(43823) (47160)66153    (87615)   (2-3) + https://github.com/lukeed/polkadot
fiber (go)(1) text        40000(162951-107000)51605.92 (72000)   60-78-109%  (1) не реплицируется
sanic -                   18258 (1)   -       45000    (57494)   26-68-95.4% (3) +
HAproxy (rust,go) (147627-110000-86000)       40746    (63860)      61-86%
 rocket (rust)(1) 30273.00(88437 - 135000)    36841    (61000)   45-55-78%   (1) text (native),   (135000/24000)
 nginx (rust,go) (1)      (37000 - 50546)     29000    (50500)   43-56-78% 
 traefik (rust,go)                            27000    (39837)      40-57%
 aiohttp                 10108.09 (1) -       30000       14-43-60%          (3) +
 express.js              7200 (1)    -        24000       10-36-51%  (6-9) запускать через pm2 start express.js -i 6
 fastrouter (go) (1)     (94831 - 153713)     16800.12 (25515)  23%        (1) не реплицируется
 echo (go)(1) text       (98912 - 160485)     14703    (24840)  22%        (1) не реплицируется
 lowhttpserver           pm2, без реплик-     20600  31-43%     (1) запускать только через pm2,  не реплицируется
 hyperexpress            pm2, без реплик-     20000  31-43%     (1) запускать только через pm2,  не реплицируется
 bun.js                  pm2, без реплик-     20000  31-43%     (1) запускать только через pm2,  не реплицируется
  rayo.js                     11912.70 (1)    15377.73 (18386)  18-23%     (1) не реплицируется
  blacksheep              5622.84 (1) -         8710.87         8-13%      (2) не реплицируется (больше 8000)
  fastapi                 4290.38(5850) (1) -   8371.09(11620)  6.4-12% (2) не реплицируется (больше 8000, больше 2)



# sanic
# http://localhost:3200/test?test=tes
# docker run --network host --rm jordi/ab -k -c 100 -n 10000 http://localhost:3200/test?test=test
# 14520.41  - 18258.40 | 404 10563
# docker run --network host --rm jordi/ab -k -c 100 -n 600000 http://localhost:3200/test?test=test
# sanic (3 * swarm) good
# 47524 - 57494 | 404 31772
# sanic (2 * swarm) ok
# 39418 | 404 22619 -
# sanic (4 * swarm)
# docker run --network host --rm jordi/ab -k -c 100 -n 600000 http://localhost:3200/test?test=test
# 50000 - 55084 | 404 41347

# docker run --network host --rm jordi/ab -k -c 100 -n 100000 http://localhost:3200/test?test=test
# docker run --network host --rm jordi/ab -k -c 100 -n 300000 http://localhost:3200/test?test=test
# 18900

# aiohttp
# 10108 - 11500 | 404 5798.53
# aiohttp (3 * swarm) good
# 29017 - 30000 | 404 14999
# aiohttp (2 * swarm) ok
# 19626 | 404 11131
# aiohttp (4 * swarm) 
# 30000 | 404 12550

# fastapi
# 5850 - | 404 6402
# fastapi (2 * swarm) good
# docker run --network host --rm jordi/ab -k -c 100 -n 200000 http://localhost:3200/test?test=test
# 10327 
# fastapi (3 * swarm) same
# 11620 | 404 12110

# fiber
# docker run --network host --rm jordi/ab -k -c 100 -n 300000 http://localhost:3000/test?test=test
# 70482.11 (68000)

# docker run --network host --rm jordi/ab -k -c 100 -n 800000 http://localhost:3000/test?test=test
# 61957.02 (66000) - 74329 # | 404 163222

# docker run --network host --rm jordi/ab -k -c 100 -n 2000000 http://localhost:3000/fweew
# 404 164162

# fast router
# 25515 | 404 153713

# echo
# 24840 | 404 160485

# actix 
# 88000 - 95381 | 404 210279

# warp
# 88000 - 99199 | simple 189000 | 404 98000

# rocket
# 55720(56000) - 61000 | simple 135000 | 404 24314

# polkadot.js >= sanic, go (x1.5    ) / fastrouter (x2)
# docker run --network host --rm jordi/ab -k -c 100 -n 500000 http://localhost:3000/test?test=test
# 43823 | 404 59704
# polka (2 * pm2) >= go - rust / rocket
# 71653 - 78000 | 404 107392
# polka (3 * pm2) = actix / rust
# docker run --network host --rm jordi/ab -k -c 100 -n 2000000 http://localhost:3000/test?test=test
# 80551  - 87615 | 404 148680 - 156074 - 
# polka (4 * pm2) = same

# rayo.js 
# 18386

# // rps
# # 1 fiber - 49957.73 / 117594.37
# # 72254 / 162951

# # 2 actix - 59381.96 / 130528.95
# # 86631 | 198159

# # 3 HAproxy balancer 39830.72 / 85283.02
# # 63860.86 | 147627


# # 3 nginx balancer 32241 / 36228.30
# # 49000 - 50546 | 59010 - 61917


# # 3 traefik balancer 27439 / 39686
# # 39837 | 58273

