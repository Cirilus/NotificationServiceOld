from aiogram.types import ReplyKeyboardMarkup

main_menu_KeyBoard = ReplyKeyboardMarkup(row_width=2, resize_keyboard=True)
main_menu_KeyBoard.add("/login")
main_menu_KeyBoard.add("/help")
