from loguru import logger

from aiogram import Dispatcher
from .handlers import cmd_start


def register(dp: Dispatcher):
    logger.info("Registering the client handlers")
    dp.register_message_handler(cmd_start, commands=['start', 'help'])
