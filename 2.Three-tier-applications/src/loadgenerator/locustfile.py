import random
from locust import HttpUser, TaskSet, between

remainders = [
    'Hola a todos!',
    'Ola pessoal!',
    'Oi gente!',
    'Hi there!',
    'Hellow everyone',
    'Buenos dias',
    'Bom dia',
    'Good morning',
    'Y que fue?']

def index(l):
    l.client.get("/")

def addRemainder(l):
    l.client.post("/add", {'remainder': random.choice(remainders)})

class UserBehavior(TaskSet):

    def on_start(self):
        index(self)

    tasks = {   index: 3,
                addRemainder: 2,
            }

class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(40, 60)
