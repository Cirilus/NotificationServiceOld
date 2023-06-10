import datetime
import os
import time

from dotenv import load_dotenv
from EmailManager import EmailSender
from postgresDAO import PostgresDAO
from loguru import logger


load_dotenv()


email = os.getenv("EMAIL")
password = os.getenv("PASSWORD")

db_user = os.getenv("DB_USER")
db_password = os.getenv("DB_PASSWORD")
db_host = os.getenv("DB_HOST")
db_port = os.getenv("DB_PORT")
db_database = os.getenv("DB_DATABASE")


sender = EmailSender(email, password)
db = PostgresDAO(db_user, db_password, db_host, db_port, db_database)

while True:
    user_mails = db.take_user_mails(datetime.date.today())
    sent_ids = []
    logger.info(f"take {len(user_mails)} for sending")
    for user_mail in user_mails:
        try:
            logger.info(f"Sending message to {user_mail[1]} with subject {user_mail[2]} and with body {user_mail[3]}")
            sender.send_message(user_mail[1], user_mail[2], user_mail[3])
            sent_ids.append(user_mail[0])
        except Exception as e:
            logger.error(f"cannot send message, err= {e}")
    try:
        db.set_sent(sent_ids)
    except Exception as e:
        logger.error(f"cannot set is_sent flag, err= {e}")
    time.sleep(30)
