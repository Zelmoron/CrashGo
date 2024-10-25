import telebot
import json
from telebot import types  # Импортируем типы для работы с кнопками

settings_file_path = 'settings.json'

# Считывание токена из файла
with open(settings_file_path, 'r') as file:
    settings = json.load(file)
    token = settings.get('token')  # Получаем значение поля "token"

telebot_TOKEN = token
telebot = telebot.TeleBot(telebot_TOKEN)

@telebot.message_handler(commands=['start'])
def send_welcome(message):
    response = "Ссылка на игру: http://t.me/CraSh_GoBot/CrashGoApp"
    
    # Создаем inline-кнопку
    keyboard = types.InlineKeyboardMarkup()
    play_button = types.InlineKeyboardButton(text="Play", url="http://t.me/CraSh_GoBot/CrashGoApp")
    keyboard.add(play_button)  # Добавляем кнопку в клавиатуру

    # Отправляем сообщение с клавиатурой
    telebot.send_message(message.chat.id, response, reply_markup=keyboard)

if __name__ == '__main__':
    telebot.polling(none_stop=True)
