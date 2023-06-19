from typing import Dict, Callable, Any, Awaitable

from aiogram import types
from aiogram.dispatcher.middlewares import BaseMiddleware
from aiogram.dispatcher.handler import CancelHandler, current_handler


class Authentication(BaseMiddleware):
    async def on_process_message(self, message: types.Message, data: dict):
        handler = current_handler.get()
        print(f"handler={handler}")
        print(f"message={message}")
        print(f"data={data}")