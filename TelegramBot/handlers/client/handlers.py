from aiogram import types
from aiogram.dispatcher import FSMContext
from loguru import logger
from KeyBoards import start as KB
from .FSMMachines import FSMAuth
from create_bot import db


async def cmd_start(message: types.Message) -> None:
    logger.debug(f"sending the start panel to the user {message.from_user.id}")
    await message.answer('Before receiving the notification, you must login', reply_markup=KB.main_menu_KeyBoard)


async def cmd_login(message: types.Message, state: FSMContext):
    async with state.proxy() as data:
        data["id"] = message.from_user.id

    await FSMAuth.email.set()
    await message.answer("Введите ваш логин")


async def login_email(message: types.Message, state: FSMContext):
    async with state.proxy() as data:
        data["email"] = message.text

    await FSMAuth.next()
    await message.answer("Введите ваш пароль")


def authenticate(user: str, password: str) -> str | None:
    return None


async def login_password(message: types.Message, state: FSMContext):
    await state.finish()
    async with state.proxy() as data:
        data["password"] = message.text
        user_id = authenticate(data["email"], data["password"])
        if user_id is None:
            await message.answer("The authentication error")
            return
        telegram_id = data["id"]
        db.set_id(user_id, telegram_id)
