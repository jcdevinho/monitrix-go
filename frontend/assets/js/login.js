export function handleLogin() {
  const loginForm = document.getElementById('loginForm')

  if (!loginForm) return

  loginForm.onsubmit = (e) => {
    e.preventDefault()
    const email = loginForm.email.value.trim()
    const password = loginForm.password.value

    if (!email || password.length < 6) {
      alert('Por favor, preencha todos os campos corretamente. A senha deve ter pelo menos 6 caracteres.')
      return
    }

    alert(`Login realizado com sucesso!\nEmail: ${email}`)
    loginForm.reset()
  }
}
