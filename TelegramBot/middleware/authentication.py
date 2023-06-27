from aiogram import types
from aiogram.dispatcher.middlewares import BaseMiddleware
from aiogram.dispatcher.handler import CancelHandler, current_handler


class Authentication(BaseMiddleware):
    async def on_process_message(self, message: types.Message, data: dict):
        pass
