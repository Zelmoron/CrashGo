import telebot
import json

# Путь к вашему файлу settings.json
settings_file_path = 'settings.json'

# Считывание токена из файла
with open(settings_file_path, 'r') as file:
    settings = json.load(file)
    token = settings.get('token')  # Получаем значение поля "token"

# Замените на токен вашего бота
BOT_TOKEN = token
bot = telebot.TeleBot(BOT_TOKEN)

@bot.message_handler(commands=['start'])
def send_welcome(message):
    response = "Ссылка на игру: http://t.me/CraSh_GoBot/CrashGoApp"
    bot.send_message(message.chat.id, response)

if __name__ == '__main__':
    bot.polling(none_stop=True)