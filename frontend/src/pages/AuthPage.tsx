import { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { GoogleLogin, CredentialResponse } from '@react-oauth/google'
import { Lock, Mail } from 'lucide-react'
import { useAuth } from '../hooks/useAuth'
import { api } from '../api'

export default function AuthPage() {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const { login } = useAuth()
  const [tab, setTab] = useState<'signin' | 'signup'>('signin')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleGoogleSuccess = async (credentialResponse: CredentialResponse) => {
    if (!credentialResponse.credential) return
    try {
      // 解析 Google ID Token（JWT）获取用户信息
      const payload = JSON.parse(atob(credentialResponse.credential.split('.')[1]))
      const { token, user } = await api.auth.googleToken({
        credential: credentialResponse.credential,
        name: payload.name,
        email: payload.email,
        picture: payload.picture,
        google_id: payload.sub,
      })
      login(token, user)
      navigate('/')
    } catch {
      setError(t('auth.googleError') || 'Google 登录失败，请重试')
    }
  }

  const handleGoogleError = () => {
    setError(t('auth.googleError') || 'Google 登录失败，请重试')
  }

  return (
    <div className="min-h-[80vh] flex">
      <div className="hidden lg:flex lg:w-1/2 bg-gradient-to-br from-gray-900 to-gray-800 text-white flex-col justify-center items-center p-12">
        <div className="max-w-md text-center">
          <div className="text-6xl mb-6">&#9889;</div>
          <h2 className="text-3xl font-bold mb-4">Skills Hub</h2>
          <p className="text-gray-300 text-lg leading-relaxed">
            A community-driven AI skills repository for discovering, creating, learning, and sharing.
          </p>
        </div>
      </div>

      <div className="flex-1 flex items-center justify-center p-8">
        <div className="w-full max-w-md">
          <div className="text-center mb-8">
            <div className="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Lock className="w-6 h-6 text-gray-600" />
            </div>
            <h1 className="text-2xl font-bold text-gray-900">{t('auth.title')}</h1>
            <p className="text-sm text-gray-500 mt-1">{t('auth.subtitle')}</p>
          </div>

          {error && (
            <div className="mb-4 p-3 bg-red-50 text-red-600 rounded-lg text-sm">{error}</div>
          )}

          {import.meta.env.VITE_GOOGLE_CLIENT_ID ? (
            <div className="flex justify-center mb-4">
              <GoogleLogin
                onSuccess={handleGoogleSuccess}
                onError={handleGoogleError}
                useOneTap
                theme="outline"
                shape="rectangular"
                size="large"
                text="signin_with"
                width="400"
              />
            </div>
          ) : (
            <div className="mb-4 p-3 bg-yellow-50 text-yellow-700 rounded-lg text-sm text-center">
              Google 登录未配置，请设置 VITE_GOOGLE_CLIENT_ID
            </div>
          )}

          <div className="relative my-6">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-gray-200" />
            </div>
            <div className="relative flex justify-center text-xs">
              <span className="px-2 bg-white text-gray-400">{t('auth.or')}</span>
            </div>
          </div>

          <div className="flex mb-6 bg-gray-100 rounded-lg p-1">
            <button onClick={() => setTab('signin')} className={`flex-1 py-2 text-sm font-medium rounded-md transition-colors $\{tab === 'signin' ? 'bg-white text-gray-900 shadow-sm' : 'text-gray-500'\}`}>
              {t('auth.signIn')}
            </button>
            <button onClick={() => setTab('signup')} className={`flex-1 py-2 text-sm font-medium rounded-md transition-colors $\{tab === 'signup' ? 'bg-white text-gray-900 shadow-sm' : 'text-gray-500'\}`}>
              {t('auth.signUp')}
            </button>
          </div>

          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('auth.email')}</label>
              <div className="relative">
                <Mail className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-gray-400" />
                <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder={t('auth.emailPlaceholder')} className="w-full pl-10 pr-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500" />
              </div>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('auth.password')}</label>
              <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder={t('auth.passwordPlaceholder')} className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500" />
            </div>
            <button className="w-full py-2.5 bg-gray-900 text-white rounded-lg text-sm font-medium hover:bg-gray-800 transition-colors">
              {tab === 'signin' ? t('auth.signIn') : t('auth.signUp')}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
