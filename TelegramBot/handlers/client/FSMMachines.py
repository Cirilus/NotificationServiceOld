from aiogram.dispatcher.filters.state import State, StatesGroup


class FSMAuth(StatesGroup):
    email = State()
    password = State()

