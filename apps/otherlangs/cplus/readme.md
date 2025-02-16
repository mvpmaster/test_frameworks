g++ test.cpp

sudo ./build.sh

https://github.com/boostorg/beast#building
https://www.boost.org/users/download/

sudo ./bootstrap.sh
sudo ./b2


json

sudo apt install libjson-c-dev libjsoncpp-dev
sudo ln -s /usr/include/jsoncpp/json/ /usr/include/json

https://github.com/stardust95/TinyCompiler/issues/2


drogon_ctl create project test


g++ main.cc -lglut -lGLU -lGL -Wall -Wextra -Werror


/usr/bin/ld: cannot find -lglut

collect2: error: ld returned 1 exit status

https://qna.habr.com/q/767901
https://stackoverflow.com/questions/17252326/cannot-find-lglut-but-freeglut3-dev-is-installed


-lGL -lGLEW -lGLU -lglut

https://gist.github.com/abdullahkady/f2782157991df652c2baee0bba05b788

https://github.com/an-tao/trantor