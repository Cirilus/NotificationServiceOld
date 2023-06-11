from aiogram import types
from aiohttp import web

from telegram.bot.create_bot import TOKEN_API


async def main(request):
    url = str(request.url)
    index = url.rfind('/')
    token = url[index + 1:]
    if token == TOKEN_API:
        update = types.Update(**await request.json())
        from telegram.bot.create_bot import dp
        await dp.process_update(update)
        return web.Response()
    else:
        return web.Response(status=403)


async def notify(request):
    url = str(request.url)
    index = url.rfind('/')
    token = url[index + 1:]
    if token == TOKEN_API:
        print("Somthing")
        return web.Response()
    else:
        return web.Response(status=403)
