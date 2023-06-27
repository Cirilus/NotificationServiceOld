from aiogram import types
from aiohttp import web

from create_bot import TOKEN_API, bot


async def main(request):
    url = str(request.url)
    index = url.rfind('/')
    token = url[index + 1:]
    if token == TOKEN_API:
        update = types.Update(**await request.json())
        from create_bot import dp
        await dp.process_update(update)
        return web.Response()
    else:
        return web.Response(status=403)


async def notify(request):
    message = await request.json()
    await bot.send_message(message["id"], f"{message['title']}\n{message['body']}")
    return web.Response()
