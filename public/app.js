function showForm(formName) {
  document.querySelectorAll('.form').forEach(form => {
    form.classList.remove('active');
  });
  document.getElementById(formName).classList.add('active');
}

// Заглушка для отправки (пока без API)
function login() {
  const email = document.getElementById('login-email').value;
  console.log('Login attempt:', email);
  alert('Вход (заглушка)');
}

function signup() {
  const name = document.getElementById('signup-name').value;
  console.log('Signup attempt:', name);
  alert('Регистрация (заглушка)');
}

function resetPassword() {
  const email = document.getElementById('reset-email').value;
  console.log('Reset password:', email);
  alert('Сброс пароля (заглушка)');
}
