
from aiohttp.web import Response, HTTPMovedPermanently, HTTPNotFound # 301
# https://docs.aiohttp.org/en/latest/web_exceptions.html

# def initRoute(router):
#     async def test(x): return Response(text='test')
#     router.add_route('*', '/test', test)

res='test'*5000

async def root_redirect(request): #video_id, m3u8_file):

    try: 
        test = request.rel_url.query.get('test', None)
        if not test: return HTTPNotFound()
    except: return HTTPNotFound()
    
    return Response(text=res, status=200)
    #return HTTPMovedPermanently(url) # редирект на основной сервер # -> Origin 


class InitRedirect:
    def __init__(self, app=None, interfaces=None): 
        #global RedirectVideoService; AnyModule = interfaces.AnyModule

        app.add_route('*', '/test', root_redirect)

        # tests
        # async def rps_test_json(x): return Response(text='test'*5000)
        # app.add_route('*', '/api/rps_test_json', rps_test_json)
        # async def rps_test_redirect(x): raise HTTPMovedPermanently('http://')
        # app.add_route('*', '/api/rps_test_redirect', rps_test_redirect)



       #\f"/video/<video_id>/<m3u8_file>", methods=["HEAD", "GET"]
