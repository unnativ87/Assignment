from flask import Flask
app = Flask(__name__)

@app.route('/ping')
def ping():
    return "Pong from service 2!"

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8002)
