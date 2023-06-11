from aiogram import types
from loguru import logger
from telegram.bot.KeyBoards import start as KB


async def cmd_start(message: types.Message) -> None:
    logger.debug(f"sending the start panel to the user {message.from_user.id}")
    await message.answer('Before receiving the notification, you must login', reply_markup=KB.main_menu_KeyBoard)