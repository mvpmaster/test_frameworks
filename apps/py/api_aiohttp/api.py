



import os, sys ; sys.path.append('./') ; sys.path.append('../')
from aiohttp.web import Application, run_app, Response

app = Application()

########################################################
from services import Interfaces



async def app_factory():
    pass
    #await Interfaces.preload_settings()
    #i = Interfaces.AnyModule
    #x=i.StartAsThread_Worker( i.AnyModule, () ) # Worker
    
    from routes import API; API(app.router, Interfaces) # подключение api
    return app
########################################################

if __name__ == '__main__': run_app(app_factory() or app, host='0.0.0.0', port=os.environ.get('PORT',3000))
