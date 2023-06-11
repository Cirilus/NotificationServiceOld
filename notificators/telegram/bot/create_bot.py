import os

from loguru import logger
from aiogram import Bot, Dispatcher, types
from aiohttp import web
from dotenv import load_dotenv

load_dotenv()

TOKEN_API = os.getenv("TELEGRAM_TOKEN")

bot = Bot(token=TOKEN_API)
Bot.set_current(bot)
dp = Dispatcher(bot)
app = web.Application()




