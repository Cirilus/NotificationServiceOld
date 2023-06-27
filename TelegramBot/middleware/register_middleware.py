from aiogram import Dispatcher
from .authentication import Authentication
from loguru import logger


def register(dp: Dispatcher):
    logger.info("Register middlewares")
    dp.middleware.setup(Authentication())
