from flask import Flask
from multiprocessing import Pool, TimeoutError
import time

app = Flask(__name__)

def timer(): # skip with multiprocessing
    time.sleep(5)
    print('TIMER COMES OUT!')

@app.route("/")
def hello():
    pool = Pool(processes=1)
    res = pool.apply_async(timer)
    return "Hello World!"

if __name__ == "__main__":
    app.run()
