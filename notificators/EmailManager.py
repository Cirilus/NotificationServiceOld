import ssl
from email.message import EmailMessage
import smtplib


class EmailSender:
    def __init__(self, email, password):
        self.email_sender = email
        self._email_password = password
        self._context = ssl.create_default_context()
        self._sender = smtplib.SMTP_SSL('smtp.gmail.com', 465, context=self._context)
        self._sender.login(self.email_sender, self._email_password)

    def send_message(self, to, subject, body):
        em = EmailMessage()
        em["From"] = self.email_sender
        em["To"] = to
        em["Subject"] = subject
        em.set_content(body)

        self._sender.sendmail(self.email_sender, to, em.as_string())

    def __del__(self):
        self._sender.close()



