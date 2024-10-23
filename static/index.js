
let tg = window.Telegram.WebApp;


let h1 = document.getElementById("login");


h1.innerHTML =tg.initDataUnsafe.user.username;





function random(){
    min = 1;
    max = 3;
    num = Math.floor(Math.random()*(max-min+1) + min);

    let number = document.getElementById("number");

    number.innerHTML = num;

    const data = {
        id:num,
        telegramid:tg.initDataUnsafe.user.id
    };


    fetch('/random', {
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

}


function getInventory(){

    const data = {
        telegramid:tg.initDataUnsafe.user.id
    };
    alert(1)
    fetch('/inventory', {
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

    
}
    