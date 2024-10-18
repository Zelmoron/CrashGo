





let tg = window.Telegram.WebApp;



const data = {
    name: tg.initDataUnsafe.user.username,
    id: tg.initDataUnsafe.user.id 
};





fetch('/api/data', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
})
.then(response => response.json())
.then(data => {
    console.log('Успех:', data);
})
.catch((error) => {
    console.error('Ошибка:', error);
});