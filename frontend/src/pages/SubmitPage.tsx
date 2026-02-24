import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { Send } from 'lucide-react'
import { api } from '../api'
import { useAuth } from '../hooks/useAuth'

const categories = [
  'Coding & Development',
  'Writing & Content',
  'Data & Analytics',
  'Design & Creative',
  'Automation',
]

const skillTypes = ['general', 'ai-powered']

export default function SubmitPage() {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const { user } = useAuth()

  const [form, setForm] = useState({
    title: '', title_zh: '',
    description: '', desc_zh: '',
    category: categories[0],
    icon: '',
    content: '', content_zh: '',
    tags: '',
    skill_type: 'general',
  })
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState('')

  if (!user) {
    return (
      <div className="max-w-2xl mx-auto px-4 py-20 text-center">
        <Send className="w-12 h-12 text-gray-300 mx-auto mb-4" />
        <p className="text-gray-500 text-lg">{t('submit.loginRequired')}</p>
      </div>
    )
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    setError('')
    try {
      await api.skills.create(form)
      navigate('/')
    } catch (err: any) {
      setError(err.message || 'Failed to submit')
    } finally {
      setSubmitting(false)
    }
  }

  const update = (field: string, value: string) =>
    setForm((prev) => ({ ...prev, [field]: value }))

  return (
    <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div className="text-center mb-10">
        <h1 className="text-3xl font-bold text-gray-900 mb-3">{t('submit.title')}</h1>
        <p className="text-gray-500">{t('submit.subtitle')}</p>
      </div>

      {error && (
        <div className="mb-6 p-3 bg-red-50 text-red-600 rounded-lg text-sm">{error}</div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.titleLabel')} *</label>
            <input
              type="text" required
              value={form.title} onChange={(e) => update('title', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.titleZHLabel')}</label>
            <input
              type="text"
              value={form.title_zh} onChange={(e) => update('title_zh', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.descLabel')} *</label>
            <textarea
              required rows={3}
              value={form.description} onChange={(e) => update('description', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.descZHLabel')}</label>
            <textarea
              rows={3}
              value={form.desc_zh} onChange={(e) => update('desc_zh', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
            />
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.categoryLabel')} *</label>
            <select
              value={form.category} onChange={(e) => update('category', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 bg-white"
            >
              {categories.map((c) => <option key={c} value={c}>{c}</option>)}
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.typeLabel')}</label>
            <select
              value={form.skill_type} onChange={(e) => update('skill_type', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 bg-white"
            >
              {skillTypes.map((st) => <option key={st} value={st}>{st}</option>)}
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.iconLabel')}</label>
            <input
              type="text" placeholder="🤖"
              value={form.icon} onChange={(e) => update('icon', e.target.value)}
              className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
            />
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.tagsLabel')}</label>
          <input
            type="text" placeholder="code, review, ai"
            value={form.tags} onChange={(e) => update('tags', e.target.value)}
            className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.contentLabel')} *</label>
          <textarea
            required rows={8}
            value={form.content} onChange={(e) => update('content', e.target.value)}
            className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm font-mono focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
            placeholder="# Skill Documentation&#10;&#10;Describe your skill in Markdown..."
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1.5">{t('submit.contentZHLabel')}</label>
          <textarea
            rows={8}
            value={form.content_zh} onChange={(e) => update('content_zh', e.target.value)}
            className="w-full px-4 py-2.5 border border-gray-200 rounded-lg text-sm font-mono focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
            placeholder="# 技能文档&#10;&#10;用 Markdown 描述你的技能..."
          />
        </div>

        <button
          type="submit"
          disabled={submitting}
          className="w-full flex items-center justify-center gap-2 py-3 bg-gray-900 text-white rounded-lg font-medium hover:bg-gray-800 disabled:opacity-50 transition-colors"
        >
          <Send className="w-4 h-4" />
          {submitting ? '...' : t('submit.submitBtn')}
        </button>
      </form>
    </div>
  )
}
