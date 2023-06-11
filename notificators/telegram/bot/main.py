import os

from aiohttp import web
from dotenv import load_dotenv
from loguru import logger
from aiogram import types

from telegram.bot.create_bot import bot, dp, app, TOKEN_API
from handlers import client, webhook


webhook_path = f'/{TOKEN_API}'

logger.level("DEBUG")


async def set_webhook():
    webhook_uri = f'https://94dc-185-119-0-221.eu.ngrok.io{webhook_path}'
    await bot.set_webhook(
        webhook_uri
    )


async def on_startup(_):
    logger.info("Setting the webhooks")
    await set_webhook()


client.register_client.register(dp)
webhook.register_webhooks.register(app)

app.on_startup.append(on_startup)
logger.info("Running the app")
web.run_app(app,
            host='0.0.0.0',
            port=8000)

