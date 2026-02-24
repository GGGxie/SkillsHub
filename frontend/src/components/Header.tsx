import { Link, useLocation } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { Globe, Menu, X } from 'lucide-react'
import { useState } from 'react'
import { useAuth } from '../hooks/useAuth'

export default function Header() {
  const { t, i18n } = useTranslation()
  const { user, logout } = useAuth()
  const location = useLocation()
  const [mobileOpen, setMobileOpen] = useState(false)

  const toggleLang = () => {
    const next = i18n.language === 'en' ? 'zh' : 'en'
    i18n.changeLanguage(next)
    localStorage.setItem('language', next)
  }

  const navItems = [
    { to: '/', label: t('nav.discover') },
    { to: '/learn', label: t('nav.learn') },
    { to: '/help', label: t('nav.help') },
  ]

  const isActive = (path: string) =>
    path === '/' ? location.pathname === '/' : location.pathname.startsWith(path)

  return (
    <header className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-gray-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center gap-8">
            <Link to="/" className="flex items-center gap-2 font-bold text-xl text-gray-900">
              <span className="text-2xl">⚡</span>
              <span>Skills Hub</span>
            </Link>

            <nav className="hidden md:flex items-center gap-1">
              {navItems.map((item) => (
                <Link
                  key={item.to}
                  to={item.to}
                  className={`px-3 py-2 rounded-lg text-sm font-medium transition-colors ${
                    isActive(item.to)
                      ? 'bg-gray-100 text-gray-900'
                      : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
                  }`}
                >
                  {item.label}
                </Link>
              ))}
            </nav>
          </div>

          <div className="hidden md:flex items-center gap-3">
            <button
              onClick={toggleLang}
              className="p-2 rounded-lg text-gray-500 hover:text-gray-700 hover:bg-gray-50 transition-colors"
              title={i18n.language === 'en' ? '切换到中文' : 'Switch to English'}
            >
              <Globe className="w-5 h-5" />
            </button>

            <Link
              to="/submit"
              className="px-4 py-2 bg-gray-900 text-white text-sm font-medium rounded-lg hover:bg-gray-800 transition-colors"
            >
              {t('nav.submit')}
            </Link>

            {user ? (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  {user.avatar ? (
                    <img src={user.avatar} alt="" className="w-8 h-8 rounded-full" />
                  ) : (
                    <div className="w-8 h-8 rounded-full bg-primary-100 text-primary-600 flex items-center justify-center text-sm font-medium">
                      {user.name[0]}
                    </div>
                  )}
                  <span className="text-sm font-medium text-gray-700">{user.name}</span>
                </div>
                <button
                  onClick={logout}
                  className="text-sm text-gray-500 hover:text-gray-700 transition-colors"
                >
                  {t('nav.logout')}
                </button>
              </div>
            ) : (
              <Link
                to="/auth"
                className="text-sm font-medium text-gray-600 hover:text-gray-900 transition-colors"
              >
                {t('nav.login')}
              </Link>
            )}
          </div>

          <button
            onClick={() => setMobileOpen(!mobileOpen)}
            className="md:hidden p-2 rounded-lg text-gray-500 hover:text-gray-700"
          >
            {mobileOpen ? <X className="w-5 h-5" /> : <Menu className="w-5 h-5" />}
          </button>
        </div>

        {mobileOpen && (
          <div className="md:hidden py-4 border-t border-gray-100">
            <nav className="flex flex-col gap-1">
              {navItems.map((item) => (
                <Link
                  key={item.to}
                  to={item.to}
                  onClick={() => setMobileOpen(false)}
                  className={`px-3 py-2 rounded-lg text-sm font-medium ${
                    isActive(item.to) ? 'bg-gray-100 text-gray-900' : 'text-gray-600'
                  }`}
                >
                  {item.label}
                </Link>
              ))}
              <Link
                to="/submit"
                onClick={() => setMobileOpen(false)}
                className="px-3 py-2 rounded-lg text-sm font-medium text-gray-600"
              >
                {t('nav.submit')}
              </Link>
              <div className="flex items-center gap-2 px-3 py-2">
                <button onClick={toggleLang} className="text-sm text-gray-500">
                  <Globe className="w-4 h-4 inline mr-1" />
                  {i18n.language === 'en' ? '中文' : 'English'}
                </button>
              </div>
              {user ? (
                <button onClick={logout} className="px-3 py-2 text-sm text-left text-gray-600">
                  {t('nav.logout')}
                </button>
              ) : (
                <Link to="/auth" onClick={() => setMobileOpen(false)} className="px-3 py-2 text-sm text-gray-600">
                  {t('nav.login')}
                </Link>
              )}
            </nav>
          </div>
        )}
      </div>
    </header>
  )
}
