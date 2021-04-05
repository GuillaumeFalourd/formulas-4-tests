#!/usr/bin/python3
import sendgrid
import os
from sendgrid.helpers.mail import *

def run():
    sg = sendgrid.SendGridAPIClient(api_key=os.environ.get('SENDGRID_API_KEY'))
    
    from_email = Email("uvogline.online@gmail.com")
    to_email = To("guillaume.falourd@gmail.com")
    subject = "Sending with SendGrid is Fun"
    content = Content("text/plain", "and easy to do anywhere, even with Python")
    
    mail = Mail(from_email, to_email, subject, content)
    
    response = sg.client.mail.send.post(request_body=mail.get())
    print(response.status_code)
    print(response.body)
    print(response.headers)
