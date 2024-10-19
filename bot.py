import telecasego
import json

# Путь к вашему файлу settings.json
settings_file_path = 'settings.json'

# Считывание токена из файла
with open(settings_file_path, 'r') as file:
    settings = json.load(file)
    token = settings.get('token') # Получаем значение поля "token"

# Замените на токен вашего бота
casego_TOKEN = token
casego = telecasego.Telecasego(casego_TOKEN)

@casego.message_handler(commands=['start'])
def send_welcome(message):
    response = "Ссылка на игру: http://t.me/CraSh_Gocasego/CrashGoApp"
    casego.send_message(message.chat.id, response)

if __name__ == '__main__':
    casego.polling(none_stop=True)
