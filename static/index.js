let tg = window.Telegram.WebApp;


let h1 = document.getElementById("login");


h1.innerHTML =tg.initDataUnsafe.user.username;