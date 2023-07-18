from aiohttp import web
from loguru import logger
from create_bot import bot, dp, app, TOKEN_API, WEBHOOK_DOMAIN
from handlers import client, webhook
import middleware

logger.level("DEBUG")


async def set_webhook():
    webhook_uri = f'{WEBHOOK_DOMAIN}/{TOKEN_API}'
    await bot.set_webhook(
        webhook_uri
    )


async def on_startup(_):
    logger.info("Setting the webhook")
    await set_webhook()


client.register_client.register(dp)
webhook.register_webhooks.register(app)
middleware.register_middleware.register(dp)

app.on_startup.append(on_startup)
logger.info("Running the app")
web.run_app(app,
            host='0.0.0.0',
            port=3000)

