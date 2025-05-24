export function handleSignup() {
  const signupForm = document.getElementById('signupForm')
  if (!signupForm) return

  signupForm.onsubmit = (e) => {
    e.preventDefault()

    const name = signupForm.name.value.trim()
    const email = signupForm.email.value.trim()
    const password = signupForm.password.value

    const nameRegex = /^[A-Za-zÀ-ÿ\s]{1,45}$/ // letras com acento e espaço
    const hasNumber = /\d/
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/ // estrutura básica
    const allowedDomains = ['gmail.com', 'hotmail.com', 'outlook.com', 'yahoo.com', 'yandex.com', 'zoho.com']
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/

    // Validação do nome
    if (!name || !nameRegex.test(name) || hasNumber.test(name)) {
      alert('Nome inválido. Use apenas letras (com acento) e espaços. Máximo de 45 caracteres. Não use números ou símbolos.')
      return
    }

    // Validação do email
    if (!emailRegex.test(email) || email.includes(' ')) {
      alert('Email inválido.')
      return
    }

    const domain = email.split('@')[1]
    const isAllowedDomain = allowedDomains.includes(domain)
    const isCorporate = domain.endsWith('.com.br') && domain.split('.')[0].length >= 10

    if (!isAllowedDomain && !isCorporate) {
      alert('Use um email válido como Gmail, Outlook, etc., ou um domínio corporativo com pelo menos 10 caracteres antes do .com.br')
      return
    }

    // Validação da senha
    if (!passwordRegex.test(password)) {
      alert('Senha fraca. Use pelo menos 8 caracteres, com 1 maiúscula, 1 minúscula, 1 número e 1 símbolo.')
      return
    }

    alert(`Registrado com sucesso!\nNome: ${name}\nEmail: ${email}`)
    signupForm.reset()
  }
}
