# модули
import models # init DB
from services.any_module import AnyModule 

# storage={}

# интерфейсы передаваемые
class Interfaces:
    models = models
    AnyModule = AnyModule

    # preload db
    @staticmethod
    async def preload_settings():
        pass
        #settings=await models.Settings.get_formated_settings()
        # AnyModule.updateSettings(settings) ; print('settings releaded:', settings) # check settings