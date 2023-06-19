from loguru import logger

from aiogram import Dispatcher

from .FSMMachines import FSMAuth
from .handlers import cmd_start, cmd_login, login_email, login_password


def register(dp: Dispatcher):
    logger.info("Registering the client handlers")
    dp.register_message_handler(cmd_start, commands=['start', 'help'])

    dp.register_message_handler(cmd_login, commands=['login'], state=None)
    dp.register_message_handler(login_email, state=FSMAuth.email)
    dp.register_message_handler(login_password, state=FSMAuth.password)