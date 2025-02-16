import os, sys, random, time, threading

SERVER_ID=os.environ.get('DEF_SERVER_ID', 's1')


class AnyModule:


    @staticmethod 
    def anyMethod(redirect_url): # video_id, m3u8_file): # "/video/1488/xcg2djHckad.m3u8"
        pass

    @staticmethod
    def StartAsThread_Worker(func, args): # def (1,)
        x = threading.Thread(target=func, args=args) ; x.start() ; return x  # x.join()

    @staticmethod
    def anyWorker():
        while True:
            time.sleep(10) # pause
            # AnyWorkerMethod()


