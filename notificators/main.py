import datetime
import os
import time

from dotenv import load_dotenv
from EmailManager import EmailSender
from TelegramManager import TelegramManager
from postgresDAO import PostgresDAO
from loguru import logger


load_dotenv()


email = os.getenv("EMAIL")
email_password = os.getenv("EMAIL_PASSWORD")

telegram_domain = os.getenv("TELEGRAM_DOMAIN")
telegram_token = os.getenv("TELEGRAM_TOKEN")
telegram_url = telegram_domain + "/" + telegram_token + "/" + "notify"

db_user = os.getenv("DB_USER")
db_password = os.getenv("DB_PASSWORD")
db_host = os.getenv("DB_HOST")
db_port = os.getenv("DB_PORT")
db_database = os.getenv("DB_DATABASE")


email_manager = EmailSender(email, email_password)
telegram_manager = TelegramManager(telegram_url, telegram_token)
db = PostgresDAO(db_user, db_password, db_host, db_port, db_database)

while True:
    user_contacts = db.take_user_contacts(datetime.date.today())
    sent_ids = []
    logger.info(f"take {len(user_contacts)} for sending")
    for user_contact in user_contacts:
        user_id = user_contact[0]
        email = user_contact[1]
        telegram = user_contact[2]
        subject = user_contact[3]
        body = user_contact[4]
        try:
            if email:
                logger.debug(f"Sending message to email {email} with subject {subject} and with body {body}")
                email_manager.send_message(email, subject, body)
            if telegram:
                logger.debug(f"Sending message to telegram user {telegram} with subject {subject} and with body {body}")
                telegram_manager.sent_notification(telegram, subject, body)
            sent_ids.append(user_id)
        except Exception as e:
            logger.error(f"cannot send message, err= {e}")
    try:
        db.set_sent(sent_ids)
    except Exception as e:
        logger.error(f"cannot set is_sent flag, err= {e}")
    time.sleep(30)
