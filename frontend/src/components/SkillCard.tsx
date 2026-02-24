import { Link } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { ThumbsUp, Eye } from 'lucide-react'
import { Skill } from '../api'

const categoryColors: Record<string, string> = {
  'Coding & Development': 'bg-blue-100 text-blue-700',
  'Writing & Content': 'bg-orange-100 text-orange-700',
  'Data & Analytics': 'bg-purple-100 text-purple-700',
  'Design & Creative': 'bg-pink-100 text-pink-700',
  'Automation': 'bg-yellow-100 text-yellow-700',
}

interface Props {
  skill: Skill
  compact?: boolean
}

export default function SkillCard({ skill, compact }: Props) {
  const { i18n, t } = useTranslation()
  const isZH = i18n.language === 'zh'
  const title = isZH && skill.title_zh ? skill.title_zh : skill.title
  const desc = isZH && skill.desc_zh ? skill.desc_zh : skill.description
  const colorClass = categoryColors[skill.category] || 'bg-gray-100 text-gray-700'

  return (
    <Link
      to={`/skill/${skill.id}`}
      className="group block bg-white rounded-xl border border-gray-100 hover:border-gray-200 hover:shadow-lg transition-all duration-300 overflow-hidden"
    >
      {!compact && skill.image && (
        <div className="aspect-video bg-gradient-to-br from-gray-50 to-gray-100 flex items-center justify-center overflow-hidden">
          <img src={skill.image} alt={title} className="w-full h-full object-cover" />
        </div>
      )}
      {!compact && !skill.image && (
        <div className="aspect-video bg-gradient-to-br from-gray-50 to-gray-100 flex items-center justify-center">
          <span className="text-5xl">{skill.icon}</span>
        </div>
      )}

      <div className="p-5">
        <div className="flex items-center gap-2 mb-3">
          <span className={`px-2.5 py-0.5 rounded-full text-xs font-medium ${colorClass}`}>
            {skill.category}
          </span>
          {skill.featured && (
            <span className="px-2.5 py-0.5 rounded-full text-xs font-medium bg-amber-100 text-amber-700">
              Featured
            </span>
          )}
          {skill.skill_type === 'ai-powered' && (
            <span className="px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-700">
              AI-Powered
            </span>
          )}
        </div>

        <h3 className="text-base font-semibold text-gray-900 group-hover:text-primary-600 transition-colors mb-2 line-clamp-1">
          {compact && <span className="mr-2">{skill.icon}</span>}
          {title}
        </h3>
        <p className="text-sm text-gray-500 line-clamp-2 leading-relaxed mb-4">{desc}</p>

        <div className="flex items-center justify-between text-sm text-gray-400">
          <div className="flex items-center gap-4">
            <span className="flex items-center gap-1">
              <ThumbsUp className="w-3.5 h-3.5" /> {skill.likes}
            </span>
            <span className="flex items-center gap-1">
              <Eye className="w-3.5 h-3.5" /> {skill.views}
            </span>
          </div>
          {skill.author_name && (
            <span className="text-xs text-gray-400">{skill.author_name}</span>
          )}
        </div>
      </div>
    </Link>
  )
}
