



let tg = window.Telegram.WebApp;



const data = {
    name: tg.initDataUnsafe.user.username,
    id: tg.initDataUnsafe.user.id ,
    coins : 100
};



fetch('/users', {
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

document.getElementById('coin').addEventListener('click', function () {
    
    this.style.transform = 'rotateY(180deg)';
    setTimeout(() => {
        window.location.href = '/api'; // Замена на нужный URL
    }, 1000);
});