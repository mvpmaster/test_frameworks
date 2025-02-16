import os, sys, random, asyncio ; sys.path.append('./') ; sys.path.append('../')
from sanic import Request, Sanic, json
from sanic.response import text, redirect

app = Sanic("Balancer") # init app

########################################################
from services import Interfaces
# подгрузка настроек
@app.before_server_start
async def setup_db(app, _): 
    pass
    #await Interfaces.preload_settings()
    #i = Interfaces.AnyModule
    #x=i.StartAsThread_Worker( i.AnyModule, () ) # Worker


from routes import API; API(app, Interfaces) # подключение api
########################################################

if __name__ == '__main__': app.run(host="0.0.0.0", # important for docker
            port=os.environ.get('API_PORT', 3000))