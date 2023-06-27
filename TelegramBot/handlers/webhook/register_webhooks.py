from aiohttp import web
from loguru import logger
from create_bot import TOKEN_API
from .handlers import main, notify


def register(app: web.Application):
    logger.info("Registering the webhooks handlers")
    app.router.add_post(f'/{TOKEN_API}', main)
    app.router.add_post(f'/{TOKEN_API}/notify', notify)