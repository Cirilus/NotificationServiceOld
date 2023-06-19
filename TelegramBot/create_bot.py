import os
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram import Bot, Dispatcher
from aiohttp import web
from dotenv import load_dotenv

load_dotenv()

TOKEN_API = os.getenv("TELEGRAM_TOKEN")

bot = Bot(token=TOKEN_API)
Bot.set_current(bot)
dp = Dispatcher(bot, storage=MemoryStorage())
Dispatcher.set_current(dp)
app = web.Application()




