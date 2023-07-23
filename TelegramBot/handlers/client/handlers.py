from aiogram import types
from aiogram.dispatcher import FSMContext
from loguru import logger
from KeyBoards import start as KB
from .FSMMachines import FSMAuth
from create_bot import db, kc
from keycloak.exceptions import KeycloakAuthenticationError


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


async def login_password(message: types.Message, state: FSMContext):
    async with state.proxy() as data:
        data["password"] = message.text
        try:
            token = kc.token(username=data["email"], password=data["password"])
        except KeycloakAuthenticationError:
            await message.answer("The authentication error")
            await state.finish()
            return
        user = kc.userinfo(token['access_token'])
        telegram_id = data["id"]
        try:
            db.set_id(user["sub"], telegram_id)
        except Exception as e:
            logger.error(f"Cannot set telegram id {telegram_id} to the user {user_id}, err={e}")
            await message.answer("The authentication error")
            await state.finish()
            return
    message.answer("You are logged in")
    await state.finish()

