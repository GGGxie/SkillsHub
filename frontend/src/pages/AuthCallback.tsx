import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../hooks/useAuth'
import { api } from '../api'

// A JWT is three base64url segments separated by dots.
const JWT_PATTERN = /^[\w-]+\.[\w-]+\.[\w-]*$/

// The backend delivers the token in the URL fragment (#token=...) so it is
// never sent to a server or recorded in access logs. A legacy ?token= query
// parameter is deliberately ignored.
function readTokenFromHash(hash: string): string | null {
  const token = new URLSearchParams(hash.replace(/^#/, '')).get('token')
  return token && JWT_PATTERN.test(token) ? token : null
}

export default function AuthCallback() {
  const navigate = useNavigate()
  const { login } = useAuth()
  // Read once; the effect below removes the hash, and StrictMode runs effects twice.
  const [token] = useState(() => readTokenFromHash(window.location.hash))

  useEffect(() => {
    // Drop the fragment from the address bar and history entry right away.
    if (window.location.hash) {
      window.history.replaceState(null, '', window.location.pathname + window.location.search)
    }

    if (!token) {
      navigate('/auth')
      return
    }
    localStorage.setItem('token', token)
    api.auth.me().then((user) => {
      login(token, user)
      navigate('/')
    }).catch(() => {
      localStorage.removeItem('token')
      navigate('/auth')
    })
  }, [token, login, navigate])

  return (
    <div className="min-h-[60vh] flex items-center justify-center">
      <div className="animate-spin w-8 h-8 border-4 border-gray-200 border-t-gray-900 rounded-full" />
    </div>
  )
}
