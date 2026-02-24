import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import ReactMarkdown from 'react-markdown'
import { ArrowLeft, ThumbsUp, Share2, Download, Calendar, Tag, User, Eye } from 'lucide-react'
import { api, Skill, Comment } from '../api'
import { useAuth } from '../hooks/useAuth'

export default function SkillDetailPage() {
  const { id } = useParams()
  const navigate = useNavigate()
  const { t, i18n } = useTranslation()
  const { user } = useAuth()
  const isZH = i18n.language === 'zh'

  const [skill, setSkill] = useState<Skill | null>(null)
  const [comments, setComments] = useState<Comment[]>([])
  const [newComment, setNewComment] = useState('')
  const [liked, setLiked] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!id) return
    setLoading(true)
    Promise.all([
      api.skills.get(id).then(setSkill),
      api.skills.comments(id).then(setComments),
    ])
      .catch(() => {})
      .finally(() => setLoading(false))
  }, [id])

  const handleLike = async () => {
    if (!id || !user) return
    const res = await api.skills.like(id)
    setLiked(res.liked)
    if (skill) {
      setSkill({ ...skill, likes: skill.likes + (res.liked ? 1 : -1) })
    }
  }

  const handleComment = async () => {
    if (!id || !newComment.trim()) return
    await api.skills.addComment(id, newComment)
    setNewComment('')
    const updated = await api.skills.comments(id)
    setComments(updated)
  }

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-12">
        <div className="animate-pulse space-y-6">
          <div className="h-8 bg-gray-100 rounded w-48" />
          <div className="h-12 bg-gray-100 rounded w-3/4" />
          <div className="h-64 bg-gray-100 rounded" />
        </div>
      </div>
    )
  }

  if (!skill) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-20 text-center">
        <p className="text-gray-400 text-lg">Skill not found</p>
      </div>
    )
  }

  const title = isZH && skill.title_zh ? skill.title_zh : skill.title
  const desc = isZH && skill.desc_zh ? skill.desc_zh : skill.description
  const content = isZH && skill.content_zh ? skill.content_zh : skill.content
  const tags = skill.tags ? skill.tags.split(',').map(t => t.trim()) : []

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <button
        onClick={() => navigate(-1)}
        className="flex items-center gap-2 text-sm text-gray-500 hover:text-gray-700 mb-6 transition-colors"
      >
        <ArrowLeft className="w-4 h-4" /> {t('skill.back')}
      </button>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Main content */}
        <div className="lg:col-span-2">
          <div className="mb-6">
            <h1 className="text-3xl font-bold text-gray-900 mb-3">
              <span className="mr-3">{skill.icon}</span>{title}
            </h1>
            <p className="text-gray-500 text-lg leading-relaxed">{desc}</p>
          </div>

          <div className="flex items-center gap-3 mb-8">
            <button
              onClick={handleLike}
              className={`flex items-center gap-2 px-4 py-2 rounded-lg border text-sm font-medium transition-colors ${
                liked
                  ? 'border-primary-200 bg-primary-50 text-primary-600'
                  : 'border-gray-200 text-gray-600 hover:bg-gray-50'
              }`}
            >
              <ThumbsUp className="w-4 h-4" /> {skill.likes}
            </button>
            <button className="flex items-center gap-2 px-4 py-2 rounded-lg border border-gray-200 text-sm font-medium text-gray-600 hover:bg-gray-50">
              <Share2 className="w-4 h-4" /> {t('skill.share')}
            </button>
          </div>

          {content && (
            <div className="prose prose-gray max-w-none mb-12">
              <ReactMarkdown
                components={{
                  code: ({ children, className, ...props }) => {
                    const isBlock = className?.includes('language-')
                    return isBlock ? (
                      <pre className="bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto">
                        <code className={className} {...props}>{children}</code>
                      </pre>
                    ) : (
                      <code className="bg-gray-100 text-gray-800 px-1.5 py-0.5 rounded text-sm" {...props}>{children}</code>
                    )
                  }
                }}
              >
                {content}
              </ReactMarkdown>
            </div>
          )}

          {/* Comments */}
          <section className="border-t border-gray-100 pt-8">
            <h2 className="text-lg font-bold text-gray-900 mb-6">
              {t('skill.comments')} ({comments.length})
            </h2>

            {user ? (
              <div className="mb-6">
                <textarea
                  value={newComment}
                  onChange={(e) => setNewComment(e.target.value)}
                  placeholder={t('skill.addComment')}
                  className="w-full p-3 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
                  rows={3}
                />
                <button
                  onClick={handleComment}
                  disabled={!newComment.trim()}
                  className="mt-2 px-4 py-2 bg-gray-900 text-white text-sm font-medium rounded-lg hover:bg-gray-800 disabled:opacity-30 transition-colors"
                >
                  {t('skill.submitComment')}
                </button>
              </div>
            ) : (
              <div className="mb-6 p-4 bg-gray-50 rounded-lg text-center">
                <p className="text-sm text-gray-500">{t('skill.loginToComment')}</p>
              </div>
            )}

            <div className="space-y-4">
              {comments.map((c) => (
                <div key={c.id} className="flex gap-3 p-4 rounded-lg bg-gray-50">
                  <div className="w-8 h-8 rounded-full bg-primary-100 text-primary-600 flex items-center justify-center text-sm font-medium flex-shrink-0">
                    {c.avatar ? (
                      <img src={c.avatar} alt="" className="w-8 h-8 rounded-full" />
                    ) : (
                      c.user_name[0]
                    )}
                  </div>
                  <div className="flex-1">
                    <div className="flex items-center gap-2 mb-1">
                      <span className="text-sm font-medium text-gray-900">{c.user_name}</span>
                      <span className="text-xs text-gray-400">
                        {new Date(c.created_at).toLocaleDateString()}
                      </span>
                    </div>
                    <p className="text-sm text-gray-600">{c.content}</p>
                  </div>
                </div>
              ))}
            </div>
          </section>
        </div>

        {/* Sidebar */}
        <div className="lg:col-span-1">
          <div className="sticky top-24 space-y-6">
            <div className="bg-gray-50 rounded-xl p-6">
              <div className="w-full aspect-square bg-gradient-to-br from-gray-100 to-gray-200 rounded-lg flex items-center justify-center mb-6">
                <span className="text-6xl">{skill.icon}</span>
              </div>

              <dl className="space-y-4">
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500 flex items-center gap-1.5"><Tag className="w-3.5 h-3.5" /> {t('skill.category')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{skill.category}</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500">{t('skill.type')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{skill.skill_type}</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500 flex items-center gap-1.5"><Calendar className="w-3.5 h-3.5" /> {t('skill.created')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{new Date(skill.created_at).toLocaleDateString()}</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500 flex items-center gap-1.5"><Calendar className="w-3.5 h-3.5" /> {t('skill.updated')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{new Date(skill.updated_at).toLocaleDateString()}</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500 flex items-center gap-1.5"><User className="w-3.5 h-3.5" /> {t('skill.author')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{skill.author_name}</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-sm text-gray-500 flex items-center gap-1.5"><Eye className="w-3.5 h-3.5" /> {t('home.views')}</dt>
                  <dd className="text-sm font-medium text-gray-900">{skill.views}</dd>
                </div>
              </dl>

              {tags.length > 0 && (
                <div className="mt-4 pt-4 border-t border-gray-200">
                  <dt className="text-sm text-gray-500 mb-2">{t('skill.tags')}</dt>
                  <div className="flex flex-wrap gap-1.5">
                    {tags.map((tag) => (
                      <span key={tag} className="px-2 py-0.5 bg-gray-200 text-gray-600 rounded text-xs">
                        {tag}
                      </span>
                    ))}
                  </div>
                </div>
              )}
            </div>

            <button className="w-full flex items-center justify-center gap-2 px-4 py-3 bg-gray-900 text-white rounded-lg font-medium hover:bg-gray-800 transition-colors">
              <Download className="w-4 h-4" /> {t('skill.download')}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
