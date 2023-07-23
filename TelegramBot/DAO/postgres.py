import psycopg2
from loguru import logger


class PostgresDao:
    def __init__(self, user, password, host, port="5432", database="postgres"):
        logger.debug(f"Connecting to db with user={user} password={password} host={host} port={port} database={database}")
        self._connection = psycopg2.connect(user=user, password=password,
                                            host=host, port=port, database=database)
        self._cursor = self._connection.cursor()
        logger.info("Successfully connected to db")

    def set_id(self, user_id, telegram_id):
        sql = "INSERT INTO account (id, telegram) VALUES (%s, %s)" \
              "ON CONFLICT (id) DO UPDATE " \
              "SET telegram = %s;"
        self._cursor.execute(sql, (user_id, telegram_id, telegram_id))
        self._connection.commit()
