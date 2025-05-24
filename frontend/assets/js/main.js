import { handleSignup } from './signup.js'
import { handleLogin } from './login.js'

// Alternância de formulários
const btnSignup = document.getElementById('btnSignup')
const btnLogin = document.getElementById('btnLogin')
const signupForm = document.getElementById('signupForm')
const loginForm = document.getElementById('loginForm')

btnSignup.onclick = () => {
  signupForm.classList.remove('hidden')
  loginForm.classList.add('hidden')

  btnSignup.classList.add('border-blue-600', 'text-blue-600')
  btnSignup.classList.remove('border-transparent', 'text-gray-600')

  btnLogin.classList.add('border-transparent', 'text-gray-600')
  btnLogin.classList.remove('border-blue-600', 'text-blue-600')
}

btnLogin.onclick = () => {
  signupForm.classList.add('hidden')
  loginForm.classList.remove('hidden')

  btnLogin.classList.add('border-blue-600', 'text-blue-600')
  btnLogin.classList.remove('border-transparent', 'text-gray-600')

  btnSignup.classList.add('border-transparent', 'text-gray-600')
  btnSignup.classList.remove('border-blue-600', 'text-blue-600')
}

// Inicializa os handlers de validação
handleSignup()
handleLogin()
