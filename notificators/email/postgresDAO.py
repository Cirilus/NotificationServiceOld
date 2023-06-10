import datetime
from typing import List

import psycopg2
from loguru import logger


class PostgresDAO:
    def __init__(self, user, password, host, port="5432", database="postgres"):
        logger.debug(f"Connecting to db with user={user} password={password} host={host} port={port} database={database}")
        self._connection = psycopg2.connect(user=user, password=password,
                                            host=host, port=port, database=database)
        self._cursor = self._connection.cursor()
        logger.info("Successfully connected to db")

    def take_user_mails(self, time):
        user_mails = []
        sql = f"SELECT n.id, COALESCE(a.email, n.email), n.title, n.body as email \
            FROM notification n \
            LEFT JOIN account_notification an on n.id = an.notification_id \
            LEFT JOIN account a on a.id = an.account_id \
            WHERE n.execution < %s and n.is_sent = false;"
        try:
            self._cursor.execute(sql, [time])
            user_mails = self._cursor.fetchall()
        except Exception as e:
            logger.error(f"Cannot take the user contacts with time {time}, err= {e}")
        return user_mails

    def set_sent(self, ids: List[str]):
        sql = "UPDATE notification SET is_sent=true WHERE id=ANY(%s::uuid[]);"
        try:
            self._cursor.execute(sql, (ids,))
            self._connection.commit()
        except Exception as e:
            logger.error(f"Cannot set the sent check in ids {ids}, err= {e}")


