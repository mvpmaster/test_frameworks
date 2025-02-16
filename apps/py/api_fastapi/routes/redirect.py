
#from fastapi import APIRouter , Depends, HTTPException, Header
from fastapi.responses import Response, RedirectResponse #, JSONResponse , HTMLResponse

# router = APIRouter( prefix="", tags=[], dependencies=[], responses={} )

res='test'*5000

async def root_redirect(test): #video_id, m3u8_file):

    if not test: return Response('not found', status_code=404) # не найдено по паттерну
    return Response(res, status_code=200)
    
    #return RedirectResponse(test, status_code=301) # редирект на основной сервер # -> Origin Server


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        #global AnyModule; AnyModule = interfaces.AnyModule

        @app.get('/test')
        async def answer(test): return await root_redirect(test)
        @app.head('/test')
        async def answer(test): return await root_redirect(test)


        # tests
        # @app.get('/api/rps_test_json')
        # async def answer(): return Response("test"*5000, media_type="application/json") 
        # @app.head('/api/rps_test_redirect')
        # async def answer(): return RedirectResponse("http://", status_code=301)
        # @app.get('/api/rps_test_redirect')
        # async def answer(): return RedirectResponse("http://", status_code=301)


       #\f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"]
