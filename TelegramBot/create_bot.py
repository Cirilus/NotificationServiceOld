import os
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram import Bot, Dispatcher
from aiohttp import web
from dotenv import load_dotenv
from DAO.postgres import PostgresDao

load_dotenv()

TOKEN_API = os.getenv("TELEGRAM_TOKEN")
db_user = os.getenv("DB_USER")
db_password = os.getenv("DB_PASSWORD")
db_host = os.getenv("DB_HOST")
db_port = os.getenv("DB_PORT")
db_database = os.getenv("DB_DATABASE")


bot = Bot(token=TOKEN_API)
db = PostgresDao(db_user, db_password, db_host, db_port, db_database)
Bot.set_current(bot)
dp = Dispatcher(bot, storage=MemoryStorage())
Dispatcher.set_current(dp)
app = web.Application()




